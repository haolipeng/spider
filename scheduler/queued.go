package scheduler

import "crawler/engine"

//requestChan和workerChan在Run函数进行初始化
type QueuedScheduler struct {
	RequestChan chan engine.Request      //chan是一个worker类型，每个worker有自己的channel
	WorkerChan  chan chan engine.Request //存储worker队列的队列
}

func (s *QueuedScheduler) GetWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.RequestChan <- r
}

//有一个worker已经准备好了，可以接收Request请求了
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.WorkerChan <- w
}

//创建worker channel和request channel,
func (s *QueuedScheduler) Run() {
	//创建workerChan和requestChan
	// Engine引擎向request队列中添加元素
	// 在创建worker工作者时，将worker对应的队列反注册给QueuedScheduler
	s.WorkerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			//request和worker事件的发生是彼此独立的
			// 所以无法判断是先获取到request事件，还是先获取到worker事件
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			//当request队列和worker队列中数量同时大于0，取第一个元素作为活跃元素
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
				//****此处如果阻塞，后续将无法收到requestChan或workerChan上的请求了
			}

			select {
			case activeWorker <- activeRequest:
				//request请求成功投递到worker通道中时，将两者从原来队列中删除
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			case rch := <-s.RequestChan:
				requestQ = append(requestQ, rch)
			case wch := <-s.WorkerChan:
				workerQ = append(workerQ, wch)
			}
		}
	}()
}
