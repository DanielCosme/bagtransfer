# Fixity
Assurance that the digital file has remained unchanged.


[] Piping
    [ ] I need a workflow
    [ ] I need an activiy
    [ ] I need a way to trigger the workflow (start)
    [ ] I need a worker
    [ ] I need a task queue for the temporal server to connect worker and workflow

[] Bussiness logic
    - Read packages into memory
        - Into our own structure.
        - Decide which type of package it is (bagit, etc...)
        - Based on the type
            - Checksum and check.

    - Directory management
    - Check sums
    - etc...

[x] I also need the packages on the project

[ ] Create a workflow that scans the directory.
    - Read the packages.
    - Dump the original on one folder.
    - Dump the results on the other folder.

- Temporal
    - Workflow observability
    - Workflow resiliency
    - Managed retries
    - ???

- Workflows orquestrate activities?
    - Bussiness Logic as code.
- Activities
    Logic that can potentially fail or not be deterministic.
    - IO

