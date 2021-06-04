# Solution for Parenthesis Problem
Problem:
Parenthesis matching for a compiler
Implement a function that checks for appropriate parenthesis matching in a string. You need to
account for all three types:{}[](). These may be nested inside each other. There will be other
characters in the text that does not need to be validated. The function returns true if the
parenthesis are appropriately matched, otherwise false.
Examples:
True:
“{}”
“{foo()}”
{foo(bar[]) baz}”
“{
(blahblahblah)[][]
Lorem{
ipsum()
}
}”
False:
“{“
“{[}]”
“{foo(}”
“foo[bar()]}”
