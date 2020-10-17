pragma solidity ^0.4.25;

contract HelloWorld {
    string s;

    constructor() public {
        s = "HelloWorld";
    }

    function say() public view returns(string) {
        return s;
    }

    function set(string _s) public {
        s = _s;
    }
}