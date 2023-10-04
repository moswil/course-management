const { Kafka } = require('kafkajs');
const protobuf = require('protobufjs');

const PROTO_FILE_PATH =
	'../course-service/internal/core/domain/event/proto/course_event.proto'; // Update with your .proto file path
const PROTO_MESSAGE_NAME = 'CourseCreatedEvent'; // Update with your Protobuf message name

const kafka = new Kafka({
	clientId: 'my-consumer',
	brokers: ['localhost:9092'], // Kafka broker address
});

const consumer = kafka.consumer({ groupId: 'my-group' });

const run = async () => {
	await consumer.connect();
	await consumer.subscribe({ topic: 'course-events', fromBeginning: true });

	await consumer.run({
		eachMessage: async ({ topic, partition, message }) => {
			try {
				const buf = message.value;

				// Load the Protobuf schema and message type
				const root = await protobuf.load(PROTO_FILE_PATH);
				const CourseCreatedEvent = root.lookupType(PROTO_MESSAGE_NAME);

				// Parse the message using Protobuf
				const parsedMessage = CourseCreatedEvent.decode(buf);

				// Access the parsed message fields
				console.log('Received event:');
				console.log('Course ID:', parsedMessage.courseId);
				console.log('Title:', parsedMessage.title);
				console.log('Instructor:', parsedMessage.instructor);
				// Access other fields as needed

				// Commit the offset to acknowledge message processing
				await consumer.commitOffsets([
					{ topic, partition, offset: message.offset },
				]);
			} catch (error) {
				console.error('Error processing Kafka message:', error);
			}
		},
	});
};

run().catch(console.error);

process.on('SIGINT', async () => {
	await consumer.disconnect();
	console.log('Kafka consumer disconnected.');
	process.exit(0);
});
