# Globire project
The Global Business Registers (globire) project, as the name hopefully reveils, offers access to various business registers around the world in a unified way. In other words: One API for access to various registers.

This repo is the Go implementation of the *Companies House URI service*

## Structure
The structure of the project is to offer libraries for the individual registers and a general API which gives access to all the libraries through a standard format. Each library uses the various options that the registers offer, which can be an API offered by the particular register, Open Data or a combination thereof.

## State of this repo

Alpha stage. Not recommended for production use.

# Todo

- [X] Company profile
- [X] Convinience functions (HasTasks, etc)
- [ ] Improve testing 
