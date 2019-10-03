import {expect} from 'chai';
import * as supertest from 'supertest';
import {CreatedNewUserData, signupNewUser} from "../utlis/auth";

const request = supertest('http://localhost:1308');

let newUserData: CreatedNewUserData;
let topicRequestData: {name: string, tracks: Array<string>, isActive: boolean} = {name: 'Tesla, Inc.', tracks: ['Tesla', 'Elon Musk'], isActive: true};

before(async () => {  
    newUserData = await signupNewUser()
});

/**
 * basic=`echo "john@example.com:secret"|tr -d '\n'|base64 -i`
 * http POST :1308/topics "Authorization:Basic ${basic}" name="Tesla Inc." tracks:='["Tesla","Elon Musk"]'
 */
it('Should POST /topics return 200 with valid topic request data', async function () {
    const res = await request
        .post('/topics')
        .set('Authorization', newUserData.jwtToken)
        .send(topicRequestData)
        .expect(200);

    validateTopic(res.body);
});

/**
 * basic=`echo "john@example.com:secret"|tr -d '\n'|base64 -i`
 * http :1308/topics "Authorization:Basic ${basic}"
 */
it('Should GET /topics return 200 with valid topics', async function () {
    const res = await request
        .get('/topics')
        .set('Authorization', newUserData.jwtToken)
        .expect(200);

    validateTopic(res.body[0]);
});

function validateTopic(topic: object) {
    expect(topic).has.property("id").greaterThan(0);
    expect(topic).has.property("name").eq(topicRequestData.name);
    expect(topic).has.property("tracks").to.eql(topicRequestData.tracks);
    expect(topic).has.property("createdAt").not.empty;
    expect(topic).has.property("isActive").eq(topicRequestData.isActive);
}