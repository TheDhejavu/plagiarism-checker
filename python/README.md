# plagiarism checker

## Introduction

Plagiarism checker is the process of locating instances of plagiarism or copyright infringement which basically means copying someone else's without referencing the work, within a document. The widespread use of computers and the advent of the Internet have made it easier to plagiarize the work of others and get away with it, but with the help of a computer-aided plagiarism detector, we can determine the plagiarism level of two documents using the Rabin Karp algorithm.

## What is Rabin Karp?

Rabin Karp algorithm is a string matching algorithm and this algorithm makes use of hash functions and the rolling hash technique, A hash function is essentially a function that map data of arbitrary size to a value of fixed size while A rolling hash allows an algorithm to calculate a hash value without having to rehash the entire string, this allows a new hash to be computed very quickly. We need to take note of two sets of data which are pattern and text, the pattern represents the string we are trying to match in the text.
The Rabin-Karp pattern matching algorithm is an improved version of the brute force approach which compares each character of pattern with each character of text.