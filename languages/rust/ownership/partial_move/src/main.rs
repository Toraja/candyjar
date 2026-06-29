#[derive(Debug)]
struct Duck {
    name: String,
    sound: String,
}

fn main() {
    let duck = Duck {
        name: String::from("Daffy"),
        sound: String::from("Quack"),
    };

    println!("Duck: {:?}", duck);

    // Move the name field out of the struct
    let name = duck.name;
    println!("Duck's name: {}", name);

    // Also the `duck` variable is no longer valid after moving the name field
    // println!("{:?}", duck); // This line would cause a compile-time error

    // Other fields of the struct can still be accessed
    let sound = duck.sound;
    println!("Duck's sound: {}", sound);
}
