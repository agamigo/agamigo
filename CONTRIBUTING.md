# Agamigo Development

<!-- vim-markdown-toc GitLab -->

* [Project Workflow](#project-workflow)
* [Gitlab](#gitlab)
	* [Gitlab Milestones](#gitlab-milestones)
	* [Gitlab Labels](#gitlab-labels)
	* [Gitlab Issues](#gitlab-issues)
		* [Umbrella Issues](#umbrella-issues)
		* [Task Issues](#task-issues)
* [Development Scopes](#development-scopes)
	* [Scope::Specification](#scopespecification)
	* [Scope::Service](#scopeservice)
	* [Scope::UX](#scopeux)
* [Project File Structure](#project-file-structure)

<!-- vim-markdown-toc -->

## Project Workflow

Our workflow, in brief, follows a cycle of development for each new release
version of Agamigo.

Here's an overview of the **Release Cycle** workflow:

1. A [**Release Milestone**](#gitlab-milestones) is created with a semantic
		version name.
	- Example: %"v0.1.0"
	- Purpose: The place to put issues that should be completed for the next
			release.
1. An [**Umbrella Issue**](#umbrella-issues) is created for each [**Dev
		Scope**](#development-scopes)
	- Example: #8
	- Purpose: Discussing/Brainstorming what should go into the next release.
1. [**Task Issues**](#task-issues) are created to track the work needed for the
		next release.
	- Example: #4
	- Purpose: Tangible tasks to be completed.
1. [**Merge Requests**](https://docs.gitlab.com/ee/user/project/merge_requests/)
		are created for [**Task Issues**](#task-issues) and both are either accepted
		or closed.
	- Example: TODO
	- Purpose: Discussion and finalization of actual changes to the project.
1. The aforementioned **Umbrella Issue** is used to determine if the project is
		ready for a new release. If so, it is closed along with the aforementioned
		**Release Milestone**.
1. **Task Issues** that remain open in the closed **Release Milestone** are
		evaluated for inclusion in the next release.

## Gitlab

This section shows our workflow from the perspective of individual GitLab
features. Handy for quick references.

### Gitlab Milestones

We currently use Milestones for:
- Release Tracking
	- Put issues that must be addressed for the next release version under the
			"vX.X.X" Milestone.

You can find the current Milestones here:
https://gitlab.com/groups/agamigo/-/milestones

**NOTE**: We only use Group level Milestones. This is because as Agamigo grows
we expect many of its components will graduate to self-contained projects under
the Agamigo Group namespace.

### Gitlab Labels

Refer to the [Agamigo Group's Labels](https://gitlab.com/groups/agamigo/-/labels)
and their descriptions for up-to-date information as to their purpose.

### Gitlab Issues

Do you have an idea, question, or bug report you'd like to share with the
project? Please feel free to open an
[issue](https://gitlab.com/agamigo/agamigo/issues) and we'll take care of
triage. :)

#### Umbrella Issues

These are issues for organizing, prioritizing, and creating [**Task
Issues**](#task-issues). Use the ~"Umbrella Issue" label.

It's a good idea to create an ~"Umbrella Issue" if you want to track work to be
done that:
- spans multiple [**Dev Scopes**](#development-scopes)
- spans multiple projects
- requires multiple people to complete
- is complex and easier to think about in phases or segments

#### Task Issues

These are issues where actual changes to the project are completed. They should
be concrete ideas and include all details of how the change should be
implemented.

These issues do not have a special label -- an issue that doesn't
have a special "* Issue" label is assumed to be a **Task Issue**.

**Task Issues** are assigned to someone who is responsible for doing the work
involved. The assignee is responsible for creating a [**Merge
Request**](https://docs.gitlab.com/ee/user/project/merge_requests/) where their
changes can be put through CI and reviewed by another project member before
being accepted.

## Development Scopes

Dev Scopes could be seen as way to determine where your skills are best offered
to the project. They are categories for Agamigo services, files, issues, ideas,
anything really.

### Scope::Specification

The **Agamigo Specification** scope consists of documentation that serves as a
reference point for internal and third-party development. This documentation
covers core data structures, APIs, protocols, and design principles that will be
used to enhance and interact with the **Agamigo Service**. As such, there is a
particular focus on stability and consensus for everything that is added to the
**Agamigo Specification**.

Dive in: ~"Scope::Specification"

### Scope::Service

The **Agamigo Service** scope consists of code and documentation that make up the
various non-user-facing software components used by Agamigo. Developers will
find their home here and will work with the **Agamigo Specification** as needed to
implement new functionality.

Dive in: ~"Scope::Service"

### Scope::UX

The **Agamigo UX** scope consists of designs, hardware, code, and documentation that
make up the various user interfaces used by Agamigo. Designers will find their
home here and will work with the **Agamigo Service** to implement new ways of
interacting with Agamigo.

Dive in: ~"Scope::UX"

## Project File Structure

- agamigo/[Design](/Design) See: [Scope::UX](#Scope::UX)
- agamigo/[Specification](/Specification) See: [Scope::Specification](#Scope::Specification)
- agamigo/[pkg](/pkg) Go packages

Please refer to the Buffalo
[documentation](https://gobuffalo.io/docs/directory-structure) for the following
directories.

- agamigo/[actions](/actions)
- agamigo/[assets](/assets)
- agamigo/[grifts](/grifts)
- agamigo/[locales](/locales)
- agamigo/[models](/models)
- agamigo/[templates](/templates)
- agamigo/[test](/test)
