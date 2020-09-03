use tonic::transport::{Channel, ClientTlsConfig};
use tonic::{metadata::MetadataValue, Request};

use dset::dset_client::DsetClient;
use dset::MetaForCaptureRequest;
use std::env;

pub mod dset {
    tonic::include_proto!("io.netograph.dset");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Setup TLS connection to grpc.netograph.io
    let tls = ClientTlsConfig::new().domain_name("grpc.netograph.io");
    let channel = Channel::from_static("https://grpc.netograph.io")
        .tls_config(tls)?
        .connect()
        .await?;

    // Netograph needs authorization info to each request
    let token = MetadataValue::from_str(&env::var("NGC_TOKEN")?)?;
    let mut client = DsetClient::with_interceptor(channel, move |mut req: Request<()>| {
        req.metadata_mut().insert("authorization", token.clone());
        Ok(req)
    });

    // craft and send the request
    let request = tonic::Request::new(MetaForCaptureRequest {
        dataset: "netograph:social".into(),
        id: "XX_VKedD3uyBsTIHqlm_Qg".into(),
        limit: 100,
        resume: "".into(),
        prefix: "".into(),
    });
    let response = client.meta_for_capture(request).await?;
    println!("> {:?}", response);

    // consume the result stream
    let mut stream = response.into_inner();
    while let Some(result) = stream.message().await? {
        println!("{:?}", result);
    }

    Ok(())
}
