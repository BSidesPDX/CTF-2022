# 200 - Khronos

## Description
A simple timing attack challenge. The user submits a flag, and it is compared character-by-character to the real one, with an added 20ms delay to reduce the number of samples needed for a good result.  The solution is in src/test/kotlin/com/github/blahblahbloopster/ApplicationTest.kt

## Deploy
```
./gradlew jar && docker build . -t khronos && docker run khronos
```
The challenge runs on port 80 (can be set in Application.kt)

## Challenge
The clock never stops.
