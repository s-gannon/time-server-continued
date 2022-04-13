# Campanile Time Server

A CS-358 project at Valparaiso University to keep accurate time for the James S. Markiewicz Solar Energy Research Facility.

![Image of clock displaying clock and solar time.](pictures/clock_preview.png)

## How to use / Installation

Please [check out the Wiki](https://github.com/nathanharmon1/time-server/wiki) for information about setting up and usage.

## Presentation and Abstract

Coming Soon

## Source Code File Structure

./clock | Files related to clock webserver and display.

./solarconversion | Files related to the conversion between clock and solar time.

./solarcli | Files to allow for command-line conversions to solar time.

./drivers | Files related to getting data off of hardware devices.

./data | The location where the software reads and writes data to.

./openntpd | Configuration files relevent to the network time protocol daemon.

./pictures | Pictures and other resources for displaying on GitHub

## Unit Testing and Linting

Automated unit test validation, compilation (where applicable), and linting is performed on the CS department's Jenkins server.
