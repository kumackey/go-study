fn plus_one(x: Option<i32>) -> Option<i32> {
    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);

    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}



fn main() {
    println!("{}", value_in_cents(Coin::Quarter(UsState::Alabama)))
}