use serde::Deserialize;
use std::error::Error;

const TOKEN_ENDPOINT: &str = "https://api.sso.cateiru.com/oauth/token";
const PUBLIC_KEY_ENDPOINT: &str = "https://api.sso.cateiru.com/oauth/jwt/key";

#[derive(Debug, Deserialize)]
pub struct PublicKey {
    pkcs8: String,
}

#[derive(Debug, Deserialize)]
pub struct TokenResponse {
    access_token: String,
    token_type: String,
    refresh_token: String,
    expires_in: String,
    id_token: String,
}

#[derive(Debug, Deserialize)]
pub struct Claims {
    name: String,
    given_name: String,
    family_name: String,
    middle_name: String,
    nick_name: String,
    preferred_username: String,
    profile: String,
    picture: String,
    website: String,
    email_verified: bool,
    gender: String,
    birthdate: String,
    zoneinfo: String,
    locale: String,
    phone_number: String,
    phone_number_verified: bool,
    updated_at: i64,

    id: String,
    role: String,
    theme: String,

    auth_time: i64,

    aud: String,
    exp: i64,
    jti: String,
    iat: i64,
    iss: String,
    nbf: i64,
    sub: String,
}

pub fn get_public_key() -> Result<String, Box<dyn Error>> {
    let response = reqwest::blocking::get(PUBLIC_KEY_ENDPOINT)?;

    let result = response.json::<PublicKey>()?;

    Ok(result.pkcs8)
}

#[cfg(test)]
mod tests {
    use crate::sso;
    use std::error::Error;

    #[test]
    fn get_public_key() -> Result<(), Box<dyn Error>> {
        let public_key = sso::get_public_key()?;

        println!("{}", public_key);

        assert!(public_key.len() != 0);

        Ok(())
    }
}
