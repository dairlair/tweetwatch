import {expect} from 'chai';
import * as supertest from 'supertest';
import {CreatedNewUserData, signupNewUser} from "../utlis/auth";

const request = supertest('http://localhost:1308');

let newUserData: CreatedNewUserData;

before(async () => {
    newUserData = await signupNewUser()
});

it('Should GET /status return 200 and JWT Token', async function () {
    const res = await request
        .get('/status')
        .set('Authorization', newUserData.jwtToken)
        .expect(200);

    expect(res.body).has.property("token").eq(newUserData.jwtToken);
    expect(res.body).has.property("id").eq(newUserData.userId);
    expect(res.body).has.property("email").eq(newUserData.email);
});