# Rust

## Requirements

- [evcxr_jupyter](https://github.com/evcxr/evcxr/blob/main/evcxr_jupyter/README.md)
  - Rust kernel for Jupyter Notebook
  - If a code block contains a reference with a non-static lifetime, code block
    (or at least the part the variable must live for) has to be surrounded.
    - Error
      ```rust
      let s = String::from("abc");
      let s1 = &s;
      println!("s: {}, s1: {}", s, s1)
      ```
    - Correct
      ```rust
      let s: String = String::from("abc");
      {
          let s1 = &s;
          println!("s: {}, s1: {}", s, s1)
      }
      ```
