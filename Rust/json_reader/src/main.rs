extern crate rustc_serialize;
use rustc_serialize::json::Json;
use std::fs::File;
use std::io::Read;
use std::env;

fn main() {

    let args: Vec<String> = env::args().collect();

    println!("arg length {}", args.len());
    
    let mut file = File::open(&args[1]).unwrap();
    let mut data = String::new();
    file.read_to_string(&mut data).unwrap();

    let json = Json::from_str(&data).unwrap();
//    println!("{}", &data[0..20]);
    println!("{}", json.find_path(&["asset", "version"]).unwrap());
}
