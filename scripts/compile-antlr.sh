#!/usr/bin/env bash

cd languages

antlr -Dlanguage=Go -visitor -listener Code.g4 -o code
antlr -Dlanguage=Go -visitor -listener Define.g4 -o define
