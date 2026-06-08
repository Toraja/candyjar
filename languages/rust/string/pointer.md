String is a fat pointer.  
The size of pointer is bigger than `usize`.

```rust
use std::mem::size_of;

println!("Size of usize: {}", size_of::<usize>());
println!("Size of string: {}", size_of::<String>());
println!("Size of &str: {}", size_of::<&str>());
```

So you can see that the pointer address of `s1` is bigger by 24 than the address of `s0`,
while the address of `s2` is bigger only by 8 than the address of `s1`.

Also, when you get the reference to a `String` as in `let s1 = &s;` below, `s1` holds the pointer to the fat pointer.  
So the the value that `s1` and `s2` holds are the same as the address of `s0` in the stack.

To get the pointer address of actual `String` value, use `as_ptr()`.

```rust
{
    let s0 = String::from("abc");
    let s1 = &s0;
    let s2 = &s0;

    println!("--- Pointer address in the stack ---");
    println!("s0: {:p}", &s0);
    println!("s1: {:p}", &s1);
    println!("s2: {:p}", &s2);

    println!("--- Actual value that each reference holds (i.e. pointer address of `s0` in the stack) ---");
    println!("s1: {:p}", s1);
    println!("s2: {:p}", s2);

    println!("--- Pointer address of actul string ---");
    println!("s0: {:?}", &s0.as_ptr());
    println!("s1: {:?}", &s1.as_ptr());
    println!("s2: {:?}", &s2.as_ptr());
}
```
