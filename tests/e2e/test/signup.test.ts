import {expect} from 'chai';
import * as supertest from 'supertest';

const request = supertest('http://localhost:1308');

var email: String, password: String;

before(async () => {  
    email = "john." + Date.now() + "@example.com";
    password = "secret";
})

it('Should POST /signup return 2000 and id and email with valid credentials', async function () {
    const res = await request
        .post('/signup')
        .send({email: email, password: password});

    expect(res.body).has.property("email").eq(email);
    expect(res.body).has.property("id").greaterThan(0);
    expect(res.body).has.property("token").not.eq("");
});

it('Should POST /signup return 422 for missing email', async function () {   
    const res = await request
        .post('/signup')
        .send({password: password})
        .expect(422);

    expect(res.body).not.has.property("token");
    expect(res.body).has.property("message").eq("email in body is required");
});

it('Should POST /signup return 422 for missing password', async function () {   
    const res = await request
        .post('/signup')
        .send({email: email})
        .expect(422);

    expect(res.body).not.has.property("token");
    expect(res.body).has.property("message").eq("password in body is required");
});

it('Should POST /signup return 422 for email already taken', async function () {   
    const res = await request
        .post('/signup')
        .send({email: email, password: password})
        .expect(422);

    expect(res.body).not.has.property("token");
    expect(res.body).has.property("message").eq("Email already taken");
});