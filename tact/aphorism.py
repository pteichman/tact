import re
import nltk
from nltk.tokenize import word_tokenize

# This is a Python port of Darius Kazemi's aphorism detector:
# http://tinysubversions.com/notes/aphorism-detection/

pronouns = re.compile("\b(i|my|me|he|she|you|his|her)\b", re.I)

def is_candidate(sentence):
    if len(sentence) < 20 or len(sentence) > 50:
        return False

    return pronouns.search(sentence) is None


def is_aphorism(sentence):
    if not is_candidate(sentence):
        return False

    tags = [v[1] for v in nltk.pos_tag(word_tokenize(sentence))]

    for i in xrange(0, len(tags)-1):
        if tags[i].startswith("NN") and tags[i+1].startswith("VB"):
            return True

    return False
