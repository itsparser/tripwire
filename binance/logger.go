package binance

import "fmt"

// Trace logs a message at level Trace on the standard logger.
func (c *Client) Trace(args ...interface{}) {
	err := c.Logger.Output(6, fmt.Sprint(args...))
	if err != nil {
		return
	}
}

// Debug logs a message at level Debug on the standard logger.
func (c *Client) Debug(args ...interface{}) {
	err := c.Logger.Output(5, fmt.Sprint(args...))
	if err != nil {
		return
	}
}

// Print logs a message at level Info on the standard logger.
func (c *Client) Print(args ...interface{}) {
	c.Logger.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func (c *Client) Info(args ...interface{}) {
	err := c.Logger.Output(4, fmt.Sprint(args...))
	if err != nil {
		return
	}
}

// Warn logs a message at level Warn on the standard logger.
func (c *Client) Warn(args ...interface{}) {
	err := c.Logger.Output(3, fmt.Sprint(args...))
	if err != nil {
		return
	}
}

// Warning logs a message at level Warn on the standard logger.
func (c *Client) Warning(args ...interface{}) {
	err := c.Logger.Output(3, fmt.Sprint(args...))
	if err != nil {
		return
	}
}

// Error logs a message at level Error on the standard logger.
//
// - Something failed but I'm not quitting.
func (c *Client) Error(args ...interface{}) {
	err := c.Logger.Output(2, fmt.Sprint(args...))
	if err != nil {
		return
	}
}

// Panic logs a message at level Panic on the standard logger.
//
// - Calls panic() after logging
func (c *Client) Panic(args ...interface{}) {
	err := c.Logger.Output(1, fmt.Sprint(args...))
	if err != nil {
		return
	}
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
//
// - Calls os.Exit(1) after logging
func (c *Client) Fatal(args ...interface{}) {
	err := c.Logger.Output(0, fmt.Sprint(args...))
	if err != nil {
		return
	}
}
