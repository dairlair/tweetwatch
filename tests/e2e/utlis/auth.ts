import * as supertest from 'supertest';

const request = supertest('http://localhost:1308');

export class CreatedNewUserData {
    public readonly email: string;
    public readonly password: string;
    public readonly jwtToken: string;
    public readonly userId: number;

    constructor(email: string, password: string, jwtToken: string, userId: number) {
        this.email = email;
        this.password = password;
        this.jwtToken = jwtToken;
        this.userId = userId;
    }
}

export async function signupNewUser() {
    let email = "john." + Date.now() + "@example.com";
    const password = "secret";
    const res = await request
        .post('/signup')
        .send({email: email, password: password});

    const userId: number = res.body.id;
    const jwtToken: string = res.body.token;

    return new CreatedNewUserData(email, password, jwtToken, userId);
}