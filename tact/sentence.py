#!/usr/bin/env python

import fileinput
import os
import nltk
import sys

from splitta import sbd 

model_path = None
for path in sys.path:
    mp = os.path.join(path, "splitta/model_svm/")
    if os.path.exists(mp):
        model_path = mp
        break

# This works like sbd.load_sbd_model, but silently.
model = sbd.SVM_Model(model_path)
model.load()

punkt = nltk.data.load("tokenizers/punkt/english.pickle")


def do_punkt(line):
    return punkt.tokenize(line.decode("utf-8", errors="replace"))


def do_splitta(line):
    return sbd.sbd_text(model, line, do_tok=False)
