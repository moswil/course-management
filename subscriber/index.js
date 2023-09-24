const { Kafka } = require('kafkajs');

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
				console.log({
					value: message.value.toString(),
					key: message.key,
					headers: message.headers,
				});
			} catch (error) {
				console.error('Error processing Kafka message:', error);
			}
		},
	});
};

run().catch(console.error);
