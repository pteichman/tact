#!/usr/bin/env python

import re
import sys


# pgstrip strips the Project Gutenberg header & footer from text.
def pgstrip(text):
    m = re.search("\*\*\* *START OF.*?\*\*\*", text)
    if m:
        text = text[m.end():]

    m = re.search("\*\*\* *END OF", text)
    if m:
        text = text[:m.start()]

    return text


def fixspace(text):
    return re.sub("\s+", " ", text.strip())


# Eliminate parenthetical statements.
def fixparens(text):
    return re.sub(" \(.*?\)", "", text)


def ispara(text):
    if len(text) < 1 or text[-1].isalnum():
        return False
    return True


def paragraphs(text):
    paras = [""]

    for line in text.split("\n"):
        line = line.rstrip()

        if line == "":
            paras.append("")
        else:
            paras[-1] += line + " "

    paras = map(fixspace, paras)
    paras = map(fixparens, paras)

    return "\n".join(filter(ispara, paras))


for file in sys.argv[1:]:
    with open(file, "r") as fd:
        print paragraphs(pgstrip(fd.read()))
