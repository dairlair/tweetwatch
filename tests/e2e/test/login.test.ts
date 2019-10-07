import {expect} from 'chai';
import * as supertest from 'supertest';
import {CreatedNewUserData, signupNewUser} from "../utlis/auth";

const request = supertest('http://localhost:1308');

let newUserData: CreatedNewUserData;

before(async () => {
    newUserData = await signupNewUser()
});

it('Should POST /login return 200 and JWT Token', async function () {
    const res = await request
        .post('/login')
        .send({email: newUserData.email, password: newUserData.password});
    console.error('Response body', JSON.stringify(res.body));
    expect(res.body).has.property("token").not.eq("");
});

it('Should POST /login return 422 for wrong credentials', async function () {
    const res = await request
        .post('/login')
        .send({email: newUserData.email, password: 'wrong password'})
        .expect(422);

    expect(res.body).not.has.property("token");
    expect(res.body).has.property("message").eq("Invalid credentials");
});