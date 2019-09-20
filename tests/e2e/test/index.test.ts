import {expect} from 'chai';
import * as supertest from 'supertest';

const request = supertest('http://localhost:1308');

describe("Index Test", () => {
    it('should always pass', function () {
        expect(true).to.equal(true);
    });
});

it('Should POST /signup return token with valid credentials', async function () {
    const email = "john." + Date.now() + "@example.com";
    const res = await request
        .post('/signup')
        .send({email: email, password: "secret"});

    expect(res.body).has.property("message");
});