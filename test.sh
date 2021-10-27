#!/bin/bash
assert() {
    expected="$1"
    input="$2"

    ./icc "$input" > tmp.s
    cc -o tmp tmp.s
    ./tmp
    actual="$?"

    if [ "$actual" = "$expected" ]; then
        echo "$input => $actual"
    else
        echo "$input => $expected expected, but got $actual"
        exit 1
    fi
}

assert 12 "5+8-1"

echo OK