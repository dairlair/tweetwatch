import {expect} from 'chai';
import * as supertest from 'supertest';
import {CreatedNewUserData, signupNewUser} from "../utlis/auth";
import { withData } from 'leche';

const request = supertest('http://localhost:1308');

type TopicRequest = {name: string, tracks: Array<string>, isActive: boolean};
type StreamRequest = {track: string};
let newUserData: CreatedNewUserData;
let topicRequestData: TopicRequest = {name: 'Tesla, Inc.', tracks: ['Tesla', 'Elon Musk'], isActive: true};

before(async () => {  
    newUserData = await signupNewUser()
});

describe('Should streams CRUD works fine', function() {
    let createdTopicId: bigint
    before(async function () {
        const res = await request
        .post('/topics')
        .set('Authorization', newUserData.jwtToken)
        .send(topicRequestData)
        .expect(200);
        expect(res.body).has.property("id").greaterThan(0)
        createdTopicId = res.body.id
    });
    const streamsToCreate = {
        simpleStream: {track: 'Something'},
        twoWordsStream: {track: 'Something else'},
    }
    withData(streamsToCreate, async function(streamRequest: StreamRequest) {
        it('Should POST /topics/:id/streams 200 with valid stream request data', async function() {
            // @TODO Add check for topic Request instanceof. When TopicRequest will be moved to separate class.
            const res = await request
                .post('/topics/' + createdTopicId + '/streams')
                .set('Authorization', newUserData.jwtToken)
                .send(streamRequest)
                .expect(200);
            validateStream(res.body, streamRequest);
        });
    });
    it('Should GET /topics/:id/streams 200 with existed streams data', async function() {
        // @TODO Add check for topic Request instanceof. When TopicRequest will be moved to separate class.
        const res = await request
            .get('/topics/' + createdTopicId + '/streams')
            .set('Authorization', newUserData.jwtToken)
            .expect(200);

        let streams = res.body
        expect(streams).length(2) // Count of data sets in streamsToCreate
    });
});

function validateStream(stream: {id: bigint, track: string}, expected: StreamRequest) {
    expect(stream).has.property("id").greaterThan(0);
    expect(stream).has.property("track").to.eql(expected.track);
    expect(stream).has.property("createdAt").not.empty;
}