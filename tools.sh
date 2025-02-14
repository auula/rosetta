#!/bin/bash

echo "WireDB Tools Shell Script."

echo ""

echo "1: testing cmd  package."
echo "2: testing conf package."
echo "3: testing vfs  package."
echo "4: testing utils package."
echo "5: testing codecov coverage."

echo ""

case_num=$1

if [ -z "$case_num" ]; then
    echo "Please provide an option (1, 2, 3, 4, 5)."
    exit 1
fi

function test_cmd_package() {
    sudo cd cmd && go test -v
}

function test_all_packages() {
    sudo go test -vet=all -race -coverprofile=coverage.out -covermode=atomic -v ./...
}

function test_utils_packages() {
    sudo cd utils && go test -v
}

function test_conf_packages(){
    cd conf && go test -v
}

function test_vfs_packages() {
    sudo cd vfs && go test -v
}

if [ "$case_num" -eq 1 ]; then
    test_cmd_package
elif [ "$case_num" -eq 2 ]; then
    test_conf_packages
elif [ "$case_num" -eq 3 ]; then
    test_vfs_packages
elif [ "$case_num" -eq 4 ]; then
    test_utils_packages
elif [ "$case_num" -eq 5 ]; then
    test_all_packages
else
    echo "Invalid option!"
    echo "Please provide a valid option (1, 2, 3, 4, 5)."
fi
