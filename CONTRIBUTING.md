# Agamigo Development

<!-- vim-markdown-toc GitLab -->

* [Gitlab Workflow](#gitlab-workflow)
	* [Gitlab Milestones](#gitlab-milestones)
	* [Gitlab Labels](#gitlab-labels)
	* [Gitlab Issues](#gitlab-issues)
* [Development Scopes](#development-scopes)
	* [Scope::Specification](#scopespecification)
	* [Scope::Service](#scopeservice)
	* [Scope::UX](#scopeux)
* [Project File Structure](#project-file-structure)

<!-- vim-markdown-toc -->

## Gitlab Workflow

This section shows what Gitlab features we use for project management, and how
they fit into our workflow.

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

Refer the [Group Labels](https://gitlab.com/groups/agamigo/-/labels) and their
descriptions for up-to-date information as to their purpose.

### Gitlab Issues

Do you have an idea, question, or bug report you'd like to share with the
project? Please feel free to open an
[issue](https://gitlab.com/agamigo/agamigo/issues)!

## Development Scopes

Dev Scopes could be seen as way to determine where your skills are best offered
to the project. They are categories for Agamigo services, files, issues, ideas,
etc.

A Gitlab label is created for each Dev Scope. Click on one to find open issues.

### Scope::Specification

The Agamigo Specification scope consists of documentation that serves as a
reference point for internal and third-party development. This documentation
covers core data structures, APIs, protocols, and design principles that will be
used to enhance and interact with the Agamigo Application. As such, there is a
particular focus on stability and consensus for everything that is added to the
Agamigo Specification.

Dive in: ~"Scope::Specification"

### Scope::Service

The Agamigo Service scope consists of code and documentation that make up the
various non-user-facing software components used by Agamigo. Developers will
find their home here and will work with the Agamigo Specification as needed to
implement new functionality.

Dive in: ~"Scope::Service"

### Scope::UX

The Agamigo UX scope consists of designs, hardware, code, and documentation that
make up the various user interfaces used by Agamigo. Designers will find their
home here and will work with the Agamigo Service to implement new ways of
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
