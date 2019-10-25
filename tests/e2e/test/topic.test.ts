import {expect} from 'chai';
import * as supertest from 'supertest';
import {CreatedNewUserData, signupNewUser} from "../utlis/auth";
import { withData } from 'leche';

const request = supertest('http://localhost:1308');

type TopicRequest = {name: string, isActive: boolean};
let newUserData: CreatedNewUserData;
let topicRequestData: TopicRequest = {name: 'Tesla, Inc.', isActive: true};
let createdTopicId: bigint;

before(async () => {  
    newUserData = await signupNewUser()
});

it('Should POST /topics return 200 with valid topic request data', async function () {
    const res = await request
        .post('/topics')
        .set('Authorization', newUserData.jwtToken)
        .send(topicRequestData)
        .expect(200);

    validateTopic(res.body, topicRequestData);
})

/**
 * basic=`echo "john@example.com:secret"|tr -d '\n'|base64 -i`
 * http :1308/topics "Authorization:Basic ${basic}"
 */
it('Should GET /topics return 200 with valid topics', async function () {
    const res = await request
        .get('/topics')
        .set('Authorization', newUserData.jwtToken)
        .expect(200);
    expect(res.body).has.a('array', 'This endpoint must returns topics list as array')
    expect(res.body).length.greaterThan(0, 'Topics list must not be empty')
    validateTopic(res.body[0], topicRequestData);
})

describe('Should topics update endpoint works fine', function() {
    before(function () {
        // @TODO Create initial topic here
    });
    withData({
        defaultTopic: topicRequestData,
        emptyTopic: {name: '', isActive: false},
        activeTopic: {name: 'Tesla, Inc.', isActive: true},
        inactiveTopic: {name: 'Tesla, Inc.', isActive: false},
    }, function(topicRequest: TopicRequest) {
        it('Should PUT /topics/:id 200 with valid topic request data', async function() {
            // @TODO Add check for topic Request instanceof. When TopicRequest will be moved to separate class.)
            const res = await request
                .put('/topics/' + createdTopicId)
                .set('Authorization', newUserData.jwtToken)
                .send(topicRequest)
                .expect(200);
            validateTopic(res.body, topicRequest);
        });
    });
});

function validateTopic(topic: {id: bigint}, expected: TopicRequest) {
    expect(topic).has.property("id").greaterThan(0);
    expect(topic).has.property("name").eq(expected.name);
    expect(topic).has.property("createdAt").not.empty;
    expect(topic).has.property("isActive").eq(expected.isActive);
    createdTopicId = topic.id;
}