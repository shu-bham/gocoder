package google

//
//import (
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestLogger_ShouldPrintMessage(t *testing.T) {
//	t.Run("should print new message", func(t *testing.T) {
//		logger := Constructor()
//		assert.Equal(t, true, logger.ShouldPrintMessage(1, "foo"), "ShouldPrintMessage() should return true for a new message")
//	})
//
//	t.Run("should not print recent message", func(t *testing.T) {
//		logger := Constructor()
//		logger.ShouldPrintMessage(1, "foo")
//		assert.Equal(t, false, logger.ShouldPrintMessage(3, "foo"), "ShouldPrintMessage() should return false for a recent message")
//	})
//
//	t.Run("should print old message again", func(t *testing.T) {
//		logger := Constructor()
//		logger.ShouldPrintMessage(1, "foo")
//		assert.Equal(t, true, logger.ShouldPrintMessage(11, "foo"), "ShouldPrintMessage() should return true for an old message")
//	})
//
//	t.Run("should handle multiple messages independently", func(t *testing.T) {
//		logger := Constructor()
//		assert.Equal(t, true, logger.ShouldPrintMessage(1, "foo"), "ShouldPrintMessage() for 'foo' at time 1 should be true")
//		assert.Equal(t, true, logger.ShouldPrintMessage(2, "bar"), "ShouldPrintMessage() for 'bar' at time 2 should be true")
//		assert.Equal(t, false, logger.ShouldPrintMessage(3, "foo"), "ShouldPrintMessage() for 'foo' at time 3 should be false")
//		assert.Equal(t, false, logger.ShouldPrintMessage(8, "bar"), "ShouldPrintMessage() for 'bar' at time 8 should be false")
//		assert.Equal(t, false, logger.ShouldPrintMessage(10, "foo"), "ShouldPrintMessage() for 'foo' at time 10 should be false")
//		assert.Equal(t, true, logger.ShouldPrintMessage(11, "foo"), "ShouldPrintMessage() for 'foo' at time 11 should be true")
//	})
//
//	t.Run("should print message after exactly 10 seconds", func(t *testing.T) {
//		logger := Constructor()
//		logger.ShouldPrintMessage(1, "foo")
//		assert.Equal(t, true, logger.ShouldPrintMessage(11, "foo"), "A message should be printable after exactly 10 seconds.")
//	})
//
//	t.Run("should handle empty message string", func(t *testing.T) {
//		logger := Constructor()
//		assert.Equal(t, true, logger.ShouldPrintMessage(1, ""), "Should be able to handle an empty message string.")
//		assert.Equal(t, false, logger.ShouldPrintMessage(5, ""), "Should not print an empty message if called within 10 seconds.")
//		assert.Equal(t, true, logger.ShouldPrintMessage(11, ""), "Should print an empty message again after 10 seconds.")
//	})
//
//	t.Run("should handle multiple interleaved messages", func(t *testing.T) {
//		logger := Constructor()
//		assert.Equal(t, true, logger.ShouldPrintMessage(0, "A"))
//		assert.Equal(t, true, logger.ShouldPrintMessage(1, "B"))
//		assert.Equal(t, true, logger.ShouldPrintMessage(2, "C"))
//		assert.Equal(t, false, logger.ShouldPrintMessage(3, "A"))
//		assert.Equal(t, false, logger.ShouldPrintMessage(4, "B"))
//		assert.Equal(t, false, logger.ShouldPrintMessage(5, "C"))
//		assert.Equal(t, true, logger.ShouldPrintMessage(10, "A"))
//		assert.Equal(t, true, logger.ShouldPrintMessage(11, "B"))
//		assert.Equal(t, true, logger.ShouldPrintMessage(12, "C"))
//	})
//
//	t.Run("should handle large timestamps", func(t *testing.T) {
//		logger := Constructor()
//		assert.Equal(t, true, logger.ShouldPrintMessage(1000000000, "foo"))
//		assert.Equal(t, false, logger.ShouldPrintMessage(1000000005, "foo"))
//		assert.Equal(t, true, logger.ShouldPrintMessage(1000000010, "foo"))
//	})
//}
