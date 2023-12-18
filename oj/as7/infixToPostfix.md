1. token between a to z add to postfix
2. token is '(' add to stack
3. token is ')' removing everything from stack and add to postfix and then remove '('
4. if len > 0 for stack and stack top element is >= i'th of string pop and add to postfix string
5. pop all reaming elements and add to postfix string
6. precedence +-, \*/, ^
