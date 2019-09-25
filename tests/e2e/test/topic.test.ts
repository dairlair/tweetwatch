import {expect} from 'chai';
import * as supertest from 'supertest';

const request = supertest('http://localhost:1308');

let email: String, password: String, userId: Number;

before(async () => {  
    email = "john." + Date.now() + "@example.com";
    password = "secret";
    const res = await request
        .post('/signup')
        .send({email: email, password: password});

    expect(res.body).has.property("id").greaterThan(0);
    expect(res.body).has.property("email").eq(email);
    userId = res.body.id;
})

/**
 * basic=`echo "john@example.com:secret"|tr -d '\n'|base64 -i`
 * http POST :1308/topics "Authorization:Basic ${basic}" name="Tesla Inc." tracks:='["Tesla","Elon Musk"]'
 */
it('Should POST /topics return 200 with valid topic request data', async function () {   
    const name = 'Tesla, Inc.';
    const tracks = ['Tesla', 'Elon Musk'];
    const buffer = Buffer.from(`${email}:${password}`)
    const res = await request
        .post('/topics')
        .set('Authorization', 'Basic ' + buffer.toString('base64'))
        .send({name: name, tracks: tracks})
        .expect(200);

    expect(res.body).has.property("id").greaterThan(0);
    expect(res.body).has.property("name").eq(name);
    expect(res.body).has.property("tracks").to.eql(tracks);
    expect(res.body).has.property("createdAt").not.empty;
    expect(res.body).has.property("isActive").eq(true);
});

/**
 * basic=`echo "john@example.com:secret"|tr -d '\n'|base64 -i`
 * http :1308/topics "Authorization:Basic ${basic}"
 */
it('Should GET /topics return 200 with valid topics', async function () {   
    const buffer = Buffer.from(`${email}:${password}`)
    const res = await request
        .get('/topics')
        .set('Authorization', 'Basic ' + buffer.toString('base64'))
        .expect(200);
});