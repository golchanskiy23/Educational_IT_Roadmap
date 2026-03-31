# Context
В первую очередь он выступает как обработчик поведения программы(допустим может закрыть ненужное соединение через timeout, что позволяет не ждать отклика зависшего сайта), к тому с помощью контекста можно переносить не требуюшие шифрования данные, чтобы например не передавать кучу аргументов в параметры функции.

## Типы context
- context.Background() Context - используется для создания корневого контекста и не может быть отменен

- TODO() Context - используется как заглушка, если вы еще не определили какой контекст вам нужен и вы его переопределите

- WithCancel(parent Context) (ctx Context, cancel CancelFunc) - создает дочерний контекст с методом отмены из родительского контекста, который может быть вызван вручную

- WithDeadline(parent Context, d time.Time) (Context, CancelFunc) - создает дочерний контекст с помощью метода отмены из родительского контекста, за исключением того, что контекст будет автоматически отменен по достижении заданного времени. Сам по себе метод Deadline возвращает время, когда задача была выполнена от имени текущего контекста.

- context.WithTimeout(parent Context, timeout time.Duration)  (Context, CancelFunc) - то же самое, что и WithDeadline, за исключением того, что он указывает время ожидания от текущего времени

- WithValue(parent Context, key, val any) Context - создает дочерний контекст из родительского контекста, который может хранить пару ключ-значение и является контекстом и также его нельзя отменить

# Context изнутри
```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
```

- Важно, что все дочерние контексты отменяются при отмене родительского.

- Важное замечание: не хранить канылы внутри структур, а передавать их как первые параметры в функции.

## Done
- Когда работа, выполняемая от имени контекста, должна быть отменена, Done возвращает канал, который закрыт

- Если этот контекст никогда не может быть отменен, Done возвращает nil

- Закрытие канала Done может произойти асинхронно после возврата функции cancel

- Последовательные вызовы Done возвращают одно и то же значение

## Err
- Если параметр Done еще не закрыт, Err возвращает значение nil

- Если параметр Done закрыт, Err возвращает ненулевую ошибку, объясняющую причину почему отменено, если контекст был отменен или отменено по истечении срока действия, если истек срок действия контекста

- После того, как Err возвращает ненулевую ошибку, последующие вызовы Err возвращают ту же ошибку

## Value
- Возвращает значение, связанное с этим контекстом для key, или nil, если с key не связано ни одно значение

- Последовательные вызовы Value с помощью одного и того же ключа возвращают один и тот же результат

- Ключ может быть любого типа, поддерживающего равенство - пакеты должны определять ключи как неэкспортируемый тип, чтобы избежать коллизий

- Ключ идентифицирует конкретное значение в контексте

- Пакеты, определяющие контекстный ключ, должны предоставлять типобезопасные средства доступа для значений, сохраненных с использованием этого ключа

Передача значений по контексту является плохой практикой. Передавайте значения по параметрам, используйте передачу по контексту только в вынужденных ситуациях. Например, можно использовать передачу по контексту логгера, какой-либо middleware айдишник, но опять же, все решения должны быть продуманы. Есть несколько причина этому, но одна из них это то, что при большой вложенности страдает перебор контекстов, также ухудшается документирование, создается проблема синхронизации работы команды.

## Внутренний тип контекста
Каждый метод контекста реализуется определенной структурой, которая реализует интерфейс Context:

### EmptyCtx
Это структура, которая никогда не отменяется, возращает nil всеми ее методами, не имеет значений и не имеет крайнего срока, также реализует интерфейс Context. Используется для создания корневого контекста.

```go
type emptyCtx struct{}

func (emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (emptyCtx) Done() <-chan struct{} {
	return nil
}

func (emptyCtx) Err() error {
	return nil
}

func (emptyCtx) Value(key any) any {
	return nil
}

type backgroundCtx struct{ emptyCtx }

func (backgroundCtx) String() string {
	return "context.Background"
}

type todoCtx struct{ emptyCtx }

func (todoCtx) String() string {
	return "context.TODO"
}
```

### CancelCtx
Это контекст, который может быть отменен, также при отмене отменяются все его дочерние элементы, которые реализуют функцию отмены. Ее создает WithCancel().

```go
type cancelCtx struct {
  	Context                        // родительский контекст

	mu       sync.Mutex            // защищает следующие поля
	done     atomic.Value          // из канала struct{}, созданной лениво, закрытым первым вызовом cancel
	children map[canceler]struct{} // устанавливается равным nil при первом отмене вызова
	err      error                 // устанавливается на отличное от nil значение при первом вызове отмены
	cause    error                 // устанавливается на отличное от nil значение при первом вызове отмены
}
```

cancelCtx определяется как контекст, который может быть отменен. Из-за древовидной структуры контекста при отмене все дочерние контексты должны быть отменены синхронно. Вам нужно просто пройти по структуре children map[canceler]structure{} и отменить их по одному.

#### Canceler
Интерфейс с методами cancel() и done(). Его имплементация - timerCtx, cancelCtx
```go
type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}

func (c *cancelCtx) Value(key any) any {
	if key == &cancelCtxKey {
		return c
	}
	return value(c.Context, key)
}

func (c *cancelCtx) Done() <-chan struct{} {
	d := c.done.Load()
	if d != nil {
		return d.(chan struct{})
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}
	return d.(chan struct{})
}

// Просто лочим мьютекс и просто возращаем саму ошибку в струкутуре контекста
func (c *cancelCtx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

type stringer interface {
	String() string
}

func contextName(c Context) string {
	if s, ok := c.(stringer); ok {
		return s.String()
	}
	return reflectlite.TypeOf(c).String()
}

func (c *cancelCtx) String() string {
	return contextName(c.Context) + ".WithCancel"
}
```

Ещё есть такая штука как PropogateCancel() - это функция, которая подписывает дочерные контексты на родительские(для мгновенной отмены выполнения операции), в целом это функция, которая обеспечивает работу дерева контекстов.

```go
func (c *cancelCtx) propagateCancel(parent Context, child canceler) {
    // назначаем родительский контекст для текущего контекста отмены
	c.Context = parent

    // если родительский контекст не поддерживает отмену (его метод Done возвращает nil), метод завершается
	done := parent.Done()
	if done == nil {
		return // родитель никогда не отменяется
	}

	select {
	case <-done:
		// родитель уже отменен
		child.cancel(false, parent.Err(), Cause(parent))
		return
	default:
	}

    // добавляем дочерний контекст в список children родительского контекста, чтобы он мог быть отменен вместе с родительским
	if p, ok := parentCancelCtx(parent); ok {
		// родитель это *cancelCtx, или является производным от него
		p.mu.Lock()
		if p.err != nil {
			// родитель уже отменен
			child.cancel(false, p.err, p.cause)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
		return
	}

	if a, ok := parent.(afterFuncer); ok {
		// родитель имплементирует AfterFunc метод
		c.mu.Lock()
		stop := a.AfterFunc(func() {
			child.cancel(false, parent.Err(), Cause(parent))
		})
		c.Context = stopCtx{
			Context: parent,
			stop:    stop,
		}
		c.mu.Unlock()
		return
	}

    // если ни один из предыдущих условий не выполнен, 
    // метод запускает новую горутину, которая ожидает сигнала об отмене 
    // родительского контекста и, в случае его получения, отменяет дочерний контекст
	goroutines.Add(1)
	go func() {
		select {
		case <-parent.Done():
			child.cancel(false, parent.Err(), Cause(parent))
		case <-child.Done():
		}
	}()
}
```

Сам cancel:
```go
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
    // просто проверяется наличие ошибок
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
  
	if cause == nil {
		cause = err
	}
  
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // уже отменен
	}
  
	c.err = err
	c.cause = cause
	d, _ := c.done.Load().(chan struct{})
	if d == nil {
		c.done.Store(closedchan)
	} else {
		close(d)
	}
  
    // отменяем все дочерние контексты
	for child := range c.children {
		child.cancel(false, err, cause)
	}
  
    // очищаем нашу мапу
	c.children = nil
	c.mu.Unlock()

    // если removeFromParent == true, то удаляем текущий контекст из дочерних
	if removeFromParent {
		removeChild(c.Context, c)
	}
}
```

#### WithCancel
WithCancel возвращает копию родителя с новым каналом Done. Канал Done контекста закрывается при вызове функции отмены или при закрытии канала Done родительского контекста. Отмена этого контекста освобождает ресурсы, поэтому код должен вызывать cancel после завершения операций в этом контексте.

- CancelFunc - обычная функция
```go
type CancelFunc func()
```

Информирует операцию о прекращении работы, но не ждет ее завершения. Ее можно вызывать одновременно несколькими подпрограммами.

```go
func withCancel(parent Context) *cancelCtx {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	c := &cancelCtx{}
	c.propagateCancel(parent, c)
	return c
}
```

### TimerCtx
Построен поверх CancelCtx, с наличием атрибута таймера и параметра дедлайна.

```go
type timerCtx struct {
	cancelCtx
	timer *time.Timer // Под cancelCtx.mu.

	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *timerCtx) String() string {
	return contextName(c.cancelCtx.Context) + ".WithDeadline(" +
		c.deadline.String() + " [" +
		time.Until(c.deadline).String() + "])"
}

func (c *timerCtx) cancel(removeFromParent bool, err, cause error) {
	c.cancelCtx.cancel(false, err, cause)
	if removeFromParent {
		// Удаляет этот timerCtx из дочерних элементов его родительского cancelCtx
		removeChild(c.cancelCtx.Context, c)
	}
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}
```

## Дополнительно:
В контексте существуют дополнительные опции, которые появились в последних версиях языка.

### WithCancelCause / Cause (Go 1.20+)

Позволяет передать произвольную причину отмены, не теряя стандартное поведение `Err()`.
```go
ctx, cancel := context.WithCancelCause(parent)
cancel(errors.New("превышен лимит запросов"))

cause := context.Cause(ctx) // "превышен лимит запросов"
err   := ctx.Err()          // context.Canceled
```

### AfterFunc
Регистрирует колбэк, который вызовется в отдельной горутине при отмене контекста.
Избавляет от ручного `select { case <-ctx.Done(): }`.
```go
stop := context.AfterFunc(ctx, func() {
    // вызовется когда ctx отменён
})
defer stop() // отменяет регистрацию если ctx ещё не отменён
```

### WithoutCancel (Go 1.21+)

Создаёт дочерний контекст, который не отменяется при отмене родителя, но наследует его значения.
Полезно для cleanup-операций после отмены основного контекста.
```go
detached := context.WithoutCancel(ctx)
```

### Ключи WithValue должны быть неэкспортируемым типом

Защита от коллизий между пакетами — два пакета могут случайно использовать одинаковые строки как ключи.
```go
// плохо
ctx = context.WithValue(ctx, "userID", 42)

// правильно
type ctxKey struct{}
ctx = context.WithValue(ctx, ctxKey{}, 42)
```

### Поиск значения идёт вверх по цепочке

`Value()` рекурсивно поднимается к родителю если ключ не найден в текущем контексте — линейный поиск O(n).
Одна из причин не злоупотреблять `WithValue` при глубокой вложенности.

## Best practices
- context.Background следует использовать только на самом высоком уровне, как корень всех производных контекстов.

- context.TODO должен использоваться, когда вы не уверены, что использовать, или если текущая функция будет использовать контекст в будущем.

- Отмены контекста рекомендуются, но эти функции могут занимать время, чтобы выполнить очистку и выход.
context.Value следует использовать как можно реже, и его нельзя применять для передачи необязательных параметров. Это делает API непонятным и может привести к ошибкам. Такие значения должны передаваться как аргументы.

- Не храните контексты в структуре, передавайте их явно в функциях, предпочтительно в качестве первого аргумента.

- Никогда не передавайте nil-контекст в качестве аргумента. Если сомневаетесь, используйте TODO.
Структура Context не имеет метода cancel, потому что только функция, которая порождает контекст, должна его отменять.

## Итог
```go
classDiagram
    class Context {
        <<interface>>
        +Deadline() (time.Time, bool)
        +Done() chan struct{}
        +Err() error
        +Value(key any) any
    }

    class emptyCtx {
        +Deadline() (time.Time, bool)
        +Done() chan struct{}
        +Err() error
        +Value(key any) any
    }

    class backgroundCtx {
        +String() string
    }

    class todoCtx {
        +String() string
    }

    class cancelCtx {
        +Context
        -mu sync.Mutex
        -done atomic.Value
        -children map[canceler]struct{}
        -err error
        -cause error
        +Done() chan struct{}
        +Err() error
        +cancel(removeFromParent bool, err, cause error)
        +propagateCancel(parent Context, child canceler)
    }

    class timerCtx {
        -timer time.Timer
        -deadline time.Time
        +Deadline() (time.Time, bool)
        +cancel(removeFromParent bool, err, cause error)
    }

    class valueCtx {
        +Context
        -key any
        -val any
        +Value(key any) any
    }

    class canceler {
        <<interface>>
        +cancel(removeFromParent bool, err, cause error)
        +Done() chan struct{}
    }

    Context <|.. emptyCtx : реализует
    Context <|.. cancelCtx : реализует
    Context <|.. timerCtx : реализует
    Context <|.. valueCtx : реализует
    emptyCtx <|-- backgroundCtx : встраивает
    emptyCtx <|-- todoCtx : встраивает
    cancelCtx <|-- timerCtx : встраивает
    canceler <|.. cancelCtx : реализует
    canceler <|.. timerCtx : реализует
```