# Gocker Sales Manual

> **Document version:** 1.0  
> **Product:** gocker v0.1.0  
> **Classification:** Internal sales enablement

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [Product Overview](#2-product-overview)
3. [Value Proposition](#3-value-proposition)
4. [Target Audience & Buyer Personas](#4-target-audience--buyer-personas)
5. [Features & Architecture](#5-features--architecture)
6. [Competitive Positioning](#6-competitive-positioning)
7. [Use Cases](#7-use-cases)
8. [Objection Handling](#8-objection-handling)
9. [Sales Scripts & Process](#9-sales-scripts--process)
10. [Technical FAQ](#10-technical-faq)
11. [Pricing Model](#11-pricing-model)
12. [Call to Action](#12-call-to-action)

---

## 1. Executive Summary

Gocker is a lightweight, open-source container runtime written in Go, designed specifically as an educational tool for developers who want to understand how Docker and Linux containers work under the hood. Unlike production-grade container runtimes that abstract away complexity, gocker exposes the core building blocks — Linux namespaces, control groups (cgroups), and process isolation — in a minimal, readable codebase.

The project occupies a unique niche at the intersection of developer education and systems programming. With fewer than 250 lines of core Go source code, gocker offers an uncluttered view into container mechanics that would otherwise require months of studying production orchestrators.

---

## 2. Product Overview

### What Is Gocker?

Gocker ("go" + "docker") is a minimal container runtime implementation that runs Linux containers using the same kernel primitives as Docker. It was created with an explicit educational mission: "The purpose of this project is to learn how Docker works and how to write a Docker by ourselves. Enjoy it, just for fun."

### Key Commands

Gocker exposes two CLI commands:

| Command | Purpose |
|---------|---------|
| `gocker run -it <command>` | Create and start a container with namespace isolation and cgroup resource limits |
| `gocker init <command>` | Internal command that initializes the container process (not intended for direct use) |

### Architecture at a Glance

```
User: gocker run -it /bin/sh
  └─ Parent process (/proc/self/exe init ...)
       └─ Creates isolated namespaces (UTS, PID, NS, NET, IPC)
            └─ Applies cgroup memory limits
                 └─ Container process executes user command
```

---

## 3. Value Proposition

### For Individual Developers

- **Learn by reading.** The entire runtime logic fits in a single afternoon of study. No sprawling codebase, no microservices architecture — just the container primitives.
- **Learn by doing.** Clone the repo, `go build`, and run your first container in minutes. Modify it, break it, fix it — the fastest path to understanding container internals.
- **Bridge the theory-practice gap.** Textbooks explain namespaces and cgroups conceptually; gocker shows you the exact system calls and filesystem interactions.

### For Teams & Organizations

- **Accelerated onboarding.** New engineers can study gocker before touching production container infrastructure, building mental models that translate directly to Docker, containerd, and Kubernetes.
- **Internal training asset.** Use gocker as a hands-on lab in container workshops, hackathons, and lunch-and-learn sessions. The low complexity means zero setup friction.
- **Auditable foundation.** Because gocker is minimal and open source, every line is auditable. Teams can verify exactly how each kernel primitive is invoked without wading through thousands of lines of abstraction.

---

## 4. Target Audience & Buyer Personas

### Persona 1: The Self-Taught Developer

- **Background:** 2-5 years experience; uses Docker daily but has never peered inside the abstraction.
- **Pain point:** "I know how to write a Dockerfile, but I don't understand how containers actually work."
- **Gocker fit:** The minimal codebase is the perfect next-step resource after Docker 101 tutorials.

### Persona 2: The Computer Science Student

- **Background:** Undergraduate or graduate student studying operating systems.
- **Pain point:** Lectures on namespaces and cgroups feel abstract without a tangible implementation.
- **Gocker fit:** A complete, runnable implementation of OS container concepts in a familiar language (Go).

### Persona 3: The Platform Engineering Team Lead

- **Background:** Leads a team building or maintaining container infrastructure.
- **Pain point:** Junior engineers lack the deep systems knowledge to debug container runtime issues.
- **Gocker fit:** Use gocker as a teaching tool in the team's onboarding curriculum and internal training workshops.

### Persona 4: The Conference Speaker / Educator

- **Background:** Creates content about Docker, containers, or Go systems programming.
- **Pain point:** Needs a clean, minimal codebase to reference in talks, blog posts, or courses.
- **Gocker fit:** gocker's brevity makes it ideal for live-coding demonstrations and annotated walkthroughs.

---

## 5. Features & Architecture

### 5.1 Linux Namespace Isolation

Gocker creates five distinct namespaces for every container, providing the isolation that makes containers feel like independent machines:

| Namespace | Flag | Purpose |
|-----------|------|---------|
| UTS | `CLONE_NEWUTS` | Isolates hostname and domain name |
| PID | `CLONE_NEWPID` | Isolates process ID numbering, giving the container its own PID 1 |
| Mount | `CLONE_NEWNS` | Isolates filesystem mount points |
| Network | `CLONE_NEWNET` | Isolates network interfaces and routing tables |
| IPC | `CLONE_NEWIPC` | Isolates inter-process communication resources |

These are applied via the `Cloneflags` field on `syscall.SysProcAttr`:

```go
cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
        syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
}
```

### 5.2 Control Groups (Cgroups) Resource Limitation

Gocker manages resource constraints through the cgroups virtual filesystem. The architecture uses a `CgroupManager` facade that delegates to individual subsystem implementations:

```
CgroupManager
  └─ SubsystemsIns[]
       └─ MemorySubSystem
            ├─ Set(path, resource)   → writes to memory.limit_in_bytes
            ├─ Apply(path, pid)      → writes PID to tasks file
            └─ Remove(path)          → deletes cgroup directory
```

**Currently implemented subsystems:**

- **Memory** (`memory.limit_in_bytes`) — limit a container's memory consumption
- The subsystem interface is designed to be extended (CPU shares, CPU sets, etc.)

### 5.3 Parent-Child Process Architecture

Gocker uses a two-phase process model:

1. **Parent process** (`Run()` in `run.go`): Forks itself via `/proc/self/exe init <command>` with namespace flags set
2. **Child process** (`init` command in `main_command.go`): Calls `container.RunContainerInitProcess()` to set up the container environment and exec into the user's command

This mirrors the architecture used by Docker (containerd → runc → container process), simplified for educational clarity.

### 5.4 Interactive Terminal Support

The `-it` flag on the `run` command maps stdin, stdout, and stderr from the host terminal directly into the container, enabling interactive shell sessions:

```bash
gocker run -it /bin/sh
```

### 5.5 Go Implementation

- **Language:** Go (single binary, no runtime dependencies)
- **CLI Framework:** urfave/cli
- **Logging:** sirupsen/logrus (structured JSON logging)
- **Dependency Management:** glide (glide.yaml / glide.lock)
- **Build:** Standard `go build` — produces a statically linked binary

---

## 6. Competitive Positioning

### Direct Competitors

| Product | Focus | Complexity | gocker Advantage |
|---------|-------|------------|------------------|
| Docker / Moby | Production container runtime | High (1M+ LOC) | Readable in an afternoon vs. weeks |
| containerd | Production container runtime | High (100k+ LOC) | Minimal surface area exposes primitives directly |
| runc | OCI runtime spec implementation | Medium (20k+ LOC) | gocker has no OCI spec overhead — pure learning |
| podman | Rootless containers, daemonless | High (100k+ LOC) | gocker focuses on one thing: "how containers work" |

### Indirect Competitors

| Product | Focus | gocker Advantage |
|---------|-------|------------------|
| Linux From Scratch | Build an entire Linux system | Container-specific — only covers namespaces/cgroups |
| Operating system textbooks | Theoretical knowledge | Runnable code, not just diagrams |
| Docker-in-Docker | Run Docker inside Docker | Teaches the runtime itself, not a setup pattern |

### gocker's Unique Position

```
Complexity
    ▲
    │           Docker / containerd / runc
    │
    │                      gocker
    │
    │           Textbook theory
    └───────────────────────────────> Educational Value
```

Gocker occupies the "Goldilocks zone" — concrete enough to run real containers, simple enough to fully understand in one sitting.

---

## 7. Use Cases

### Use Case 1: Self-Paced Learning

A developer with Docker experience wants to understand container internals.

```
┌──────────────────────────────────────────────┐
│ 1. Clone: git clone https://github.com/      │
│             inputx/gocker.git                 │
│ 2. Build: go build                           │
│ 3. Explore: read main.go, run.go,            │
│             container_process.go              │
│ 4. Run: ./gocker run -it /bin/sh              │
│ 5. Modify: add a new namespace, observe       │
│             the behavior change               │
└──────────────────────────────────────────────┘
```

### Use Case 2: Workshop / Bootcamp Lab

An instructor runs a 90-minute container internals workshop.

1. **Lecture (20 min):** Namespaces and cgroups concepts
2. **Code Walkthrough (25 min):** Instructor projects gocker source on screen
3. **Hands-On Lab (30 min):** Students run containers, inspect `/proc/<pid>/ns`, modify cgroup limits
4. **Extension Exercise (15 min):** Students add a CPU cgroup subsystem

### Use Case 3: Technical Interview Preparation

A candidate preparing for a platform engineering role uses gocker to:

- Trace the exact syscalls (`clone()`, `setns()`, cgroup filesystem writes) used in container creation
- Understand the `/proc/self/exe` re-execution pattern
- Explain how `CLONE_NEWPID` gives the container its own PID namespace

### Use Case 4: Custom Tool Prototyping

A developer building a lightweight container solution for an embedded Linux project:

- Forks gocker to strip out unneeded features
- Adds custom namespace configurations
- Ships a sub-1 MB binary with exactly the isolation needed

---

## 8. Objection Handling

### Objection 1: "It's too minimal. I can't use this in production."

**Response:** "You're right — gocker is not designed for production. That's the point. It's an educational tool, like a flight simulator for pilots. You wouldn't fly passengers in a simulator, but you'd never let a pilot fly without one. Gocker fills the same role for container infrastructure."

### Objection 2: "I can learn the same thing from reading Docker's source code."

**Response:** "Docker's source code spans thousands of files and over a million lines of code across multiple repositories (Moby, containerd, runc, etc.). The container creation logic is buried under layers of API handling, image management, networking, storage drivers, and orchestration. Gocker isolates just the container runtime core — namespace setup, cgroup configuration, and process lifecycle — in a handful of files. You can read the entire runtime in the same time it takes to find the right file in the Docker codebase."

### Objection 3: "I already understand containers theoretically."

**Response:** "Theory is a great foundation, but there's a difference between knowing that `clone()` with `CLONE_NEWPID` creates a PID namespace and seeing the exact three lines of Go code that do it, running it, and observing the result. Gocker bridges the gap from 'I understand the concept' to 'I can implement the concept.'"

### Objection 4: "Isn't this project abandoned? It has minimal activity."

**Response:** "Gocker is feature-complete for its educational mission. Adding more features would undermine its purpose — the entire value is in the minimal surface area. A stable, unchanging codebase is actually an asset for learners who want a reliable reference."

### Objection 5: "Why should my team invest time in this instead of just using Docker?"

**Response:** "Your team should use Docker for daily work. Gocker is for the one or two afternoons when they ask 'how does Docker actually do that?' The time investment is minimal — a single afternoon of study — and the payoff is deeper debugging skills, better incident response, and more confident infrastructure decisions."

---

## 9. Sales Scripts & Process

### Discovery Questions

Ask these questions to qualify a prospect:

1. "Does your team use Docker or Kubernetes in production?"
2. "Have any of your engineers ever struggled to debug a container-related issue because they didn't understand the underlying mechanics?"
3. "Do you run internal training or onboarding programs for new infrastructure engineers?"
4. "How do your junior engineers currently learn about container internals?"

### Outreach Template (Email / Slack)

```
Subject: Quick question about container knowledge on your team

Hi [Name],

I noticed your team runs [Docker / Kubernetes] in production. A common
challenge we hear is that engineers know *how* to use containers but don't
fully understand *how they work under the hood* — which makes debugging
tough when things go wrong.

Gocker is an open-source tool (github.com/inputx/gocker) that solves this
by providing a minimal container runtime — about 200 lines of Go —
that anyone can read, run, and modify in an afternoon.

Would you be open to a 15-minute call where I show how teams use gocker
as a training tool? No commitment — just a demo of what it does.

Best,
[Name]
```

### Demo Script (5 Minutes)

1. **Clone + Build (30 seconds)**
   ```
   git clone https://github.com/inputx/gocker.git
   cd gocker
   go build
   ```

2. **Run a container (30 seconds)**
   ```
   ./gocker run -it /bin/sh
   ```

3. **Show isolation (1 minute)**
   Inside container: `ps aux` (only PID 1 visible)
   Outside container: `ls /proc/<pid>/ns` (show namespace IDs)

4. **Walk through source (2 minutes)**
   Open `container_process.go` — point to the `Cloneflags` line
   Open `cgroups/cgroup_manager.go` — point to the Set/Apply/Destroy cycle

5. **Compare to Docker (1 minute)**
   "Docker does the same thing, plus a thousand other things. Gocker is the
   100% pure version of that one thing."

---

## 10. Technical FAQ

### Q: What Linux kernel version do I need?

A: Any modern Linux kernel (3.8+) with namespaces and cgroups support enabled. Most distributions ship with these features enabled by default.

### Q: Do I need root access?

A: Yes. Creating namespaces via `clone()` and writing to cgroup files requires root privileges. Run gocker with `sudo` or as root.

### Q: How do I add CPU or I/O cgroup limits?

A: Implement the `Subsystem` interface (defined in `cgroups/subsystems/subsystem.go`). The interface requires four methods: `Name()`, `Set()`, `Apply()`, and `Remove()`. Register your new subsystem in the `SubsystemsIns` slice. The memory subsystem in `memory.go` is a complete reference implementation.

### Q: Does gocker support container images?

A: No. Gocker runs a single command inside an isolated environment using the host's filesystem. Image management (layers, registries, pull/push) is intentionally out of scope. This keeps the codebase focused on the runtime primitives.

### Q: Does gocker support networking?

A: Gocker creates a network namespace (`CLONE_NEWNET`) to isolate the container's network stack, but does not configure virtual Ethernet pairs, bridges, or port forwarding. The container starts with only a loopback interface. Network configuration is left as an exercise for the learner.

### Q: Can I use gocker in a CI/CD pipeline?

A: Technically possible, but not recommended. Gocker's value is educational, not operational. Use Docker or containerd for CI pipelines.

### Q: How is gocker different from Docker?

A:

| Aspect | Docker | gocker |
|--------|--------|--------|
| Lines of code (core) | 1,000,000+ | ~200 |
| Image management | Yes | No |
| Network configuration | Yes | No |
| Volume management | Yes | No |
| Registry integration | Yes | No |
| Orchestration | Swarm/Kubernetes | No |
| Educational purpose | Incidental | Primary |
| Time to read entire runtime | Weeks | 1 hour |

### Q: What problem does `gocker init` solve?

A: The `init` command is the entry point for the container's child process. When gocker's parent process forks via `/proc/self/exe init <command>`, it re-executes the gocker binary with the `init` subcommand. This second invocation sets up the container environment and calls `exec` to replace itself with the user's command. This pattern — re-executing the runtime binary — is the same approach used by runc and containerd.

---

## 11. Pricing Model

Gocker is **free and open source** under the MIT License.

| Tier | Price | What You Get |
|------|-------|-------------|
| Open Source | Free | Full source code, documentation, issue tracker |
| — | — | — |

There is no paid tier, no enterprise license, and no premium support. The project is maintained as a community resource for the developer education ecosystem.

**Contributions welcome.** The best way to "pay" for gocker is to contribute improvements, write tutorials, give conference talks referencing it, or mentor a junior developer through their first container deep-dive.

---

## 12. Call to Action

### For Developers

```
git clone https://github.com/inputx/gocker.git
cd gocker
go build
sudo ./gocker run -it /bin/sh
```

Read the source. Modify the source. Break things. Fix them. You'll walk away understanding containers at a deeper level than 90% of Docker users.

### For Team Leads & Managers

Schedule a 90-minute container internals workshop using gocker as the lab material. Estimated outcome: every engineer walks away able to explain how namespaces, cgroups, and container processes work — with a mental model that transfers directly to Docker and Kubernetes debugging.

### For Educators

Add gocker to your course syllabus, conference talk, or blog post. The codebase is small enough to print on a few slides, complete enough to run live, and clean enough to use as a reference architecture.

---

> **Gocker: The shortest path from "I use containers" to "I understand containers."**
> [https://github.com/inputx/gocker](https://github.com/inputx/gocker)
 No newline at end of file