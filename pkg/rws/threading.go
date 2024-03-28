package rws

type TaskRunner struct {
	limitChan chan struct{}
}

func NewTaskRunner(concurrentCount int) *TaskRunner {
	return &TaskRunner{limitChan: make(chan struct{}, concurrentCount)}
}

func (re TaskRunner) Schedule(task func()) {
	re.limitChan <- struct{}{}

	go func() {
		defer func() {
			<-re.limitChan
		}()
		task()
	}()
}
