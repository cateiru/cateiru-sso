use urlencoding::encode;

const SSO_ENDPOINT: &str = "https://sso.cateiru.com/sso/login";

pub fn create_uri(client_id: &str, redirect: &str) -> String {
    let encoded_client_id = encode(client_id);
    let encoded_redirect = encode(redirect);

    return format!(
        "{}?scope=openid&response_type=code&client_id={}&redirect_uri={}&prompt=consent",
        SSO_ENDPOINT, encoded_client_id, encoded_redirect
    );
}

#[cfg(test)]
mod tests {
    use crate::client;
    use urlencoding::encode;

    #[test]
    fn create_uri() {
        let client_id = "hoge";
        let redirect = "https://example.com";

        let uri = client::create_uri(client_id, redirect);

        let encoded_client_id = encode(client_id);
        let encoded_redirect = encode(redirect);

        assert_eq!(uri, format!(
            "https://sso.cateiru.com/sso/login?scope=openid&response_type=code&client_id={}&redirect_uri={}&prompt=consent",
            encoded_client_id, encoded_redirect
        ))
    }
}
