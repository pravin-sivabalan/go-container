add to container
- [] create container stats
- [] container networking
- [] isolate env variables per process

# Specs
* Isolate env variables per container
* Support container to container communication via shared network
* Get container stats
* Get historical container stats


# What I learned...

### Linux Concepts
namespaces: controls what you can see
cgroups: controls what you can do


### Networking


### Useful commands
lsns: List of system namespaces
id

# Learn more about...

### MacOS vs Linux
When I was trying to setup containers on MacOS, I quickly learned that namespaces and cgroups were not represented the same way in MacOS as they are in Linux. But I've run Docker on my MacOS machine before, so how does that work?

* docker for macos uses HyperKit, which uses HyperVisor



### Running list of questions
User namespaces
- What is the use case of user namespaces?
- What is id mapping (UID & GID)?
- How does the the root user work?
- What does this command do (/lib/systemd/systemd --user)?
- learn more about mounts

What is docker0 (how is this implemented on MacOS)?

https://collabnix.com/how-docker-for-mac-works-under-the-hood/#:~:text=Docker%20for%20Mac%20uses%20HyperKit,built%20on%20top%20of%20Hypervisor.&text=Docker%20for%20Mac%20does%20not,%2Fvar%2Frun%2Fdocker.