fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::compile_protos("../proto/ngapi/dsetapi/dset.proto")?;
    tonic_build::compile_protos("../proto/ngapi/userapi/user.proto")?;
    Ok(())
}
