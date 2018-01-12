extern crate serde_json;

use serde_json::Value;
use std::fs::File;
use std::io::Read;
use std::env;
/*
struct BufferView {
    buffer: u32,
    byteOffset: u32,
    byteLength: u32,
}
*/
fn main() {

    let args: Vec<String> = env::args().collect();

    println!("arg length {}", args.len());
    
    let mut file = File::open(&args[1]).unwrap();
    let mut data = String::new();
    file.read_to_string(&mut data).unwrap();

    let json : Value = serde_json::from_str(&data).unwrap();
//    println!("{}", &data[0..20]);
    println!("Value: {}", json["bufferViews"][0])
}
