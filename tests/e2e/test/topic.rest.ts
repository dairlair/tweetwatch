import {expect} from 'chai';
import * as supertest from 'supertest';

const request = supertest('http://localhost:1308');

var email: String, password: String;

before(async () => {  
    email = "john." + Date.now() + "@example.com";
    password = "secret";
    const res = await request
        .post('/signup')
        .send({email: email, password: password});

    expect(res.body).has.property("id").greaterThan(0);
    expect(res.body).has.property("email").eq(email);
})

/**
 * basic=`echo "john@example.com:secret"|tr -d '\n'|base64 -i`
 * http post :1308/topic "Authorization:Basic ${basic}" name=MyTopic track=Tesla
 */
it('Should POST /topic return 200 with valid topic request data', async function () {   
    const name = 'Topic #1';
    const track = 'Tesla';
    const res = await request
        .post('/topic')
        .send({name: name, track: track})
        .expect(200);

    expect(res.body).has.property("id").greaterThan(0);
});