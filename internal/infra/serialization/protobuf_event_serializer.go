package serialization

import (
	"github.com/moswil/course-management/internal/core/domain/event/pb"
	"google.golang.org/protobuf/proto"
)

// ProtobufEventSerializer is responsible for serializing events to Protobuf format.
type ProtobufEventSerializer struct {
	// Add any necessary code here
}

// NewProtobufEventSerializer creates a new instance of ProtobufEventSerializer.
func NewProtobufEventSerializer() *ProtobufEventSerializer {
	return &ProtobufEventSerializer{}
}

// SerializeCourseCreatedEvent serializes a CourseCreatedEvent to Protobuf format.
func (s *ProtobufEventSerializer) SerializeCourseCreatedEvent(event *pb.CourseCreatedEvent) ([]byte, error) {
	pbEvent := &pb.CourseCreatedEvent{
		CourseId:   event.CourseId,
		Title:      event.Title,
		Instructor: event.Instructor,
		// Map other fields as needed
	}

	return proto.Marshal(pbEvent)
}
