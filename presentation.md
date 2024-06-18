---
title: Automated Program Repair
author: Martin Monperrus
---

# Automated Program Repair

- APR is highly dependent on unit tests.
- Unit tests should/have to cover all the edge cases.
- This is the basic type or APR (or automated big fixing) which has the unit
tests as a starting point.
- Quality tests = quality patches (fixes).
- We need a ranker, a sort of automated process (or human) that tells if the
code has a bug or not after applying the APR process. 
- 2 types of rankers:
    1. Failing ranker (checks for the assertions from unit tests).
    2. Regression ranker.
- Problem of finding the best patch (patch ranking)
