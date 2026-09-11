class StateManager<T> {
  private item: T;
  private subscribers: Set<(v: T) => void>
  constructor(value: T) {
    this.item = value;
    this.subscribers = new Set()
  }

  get() {
    return this.item
  }

  subscribe(callback: (v: T) => void) {
    this.subscribers.add(callback)
    return callback
  }

  unsubscribe(callback: (v: T) => void) {
    this.subscribers.delete(callback)
  }

  private callSubscibers() {
    this.subscribers.forEach((item) => {
      item(this.item)
    })
  }

  set(value: T) {
    this.item = value;
    this.callSubscibers()
  }

  update(mutator: (v: T) => T) {
    this.item = mutator(this.item)
    this.callSubscibers()
  }
}

const $state = <T>(item: T) => {
  return new StateManager(item)
}

console.log(`C`)
const counter = $state(2)
const item1 = counter.subscribe((v) => {
  console.log(v)
})
counter.subscribe((v) => {
  console.log(v)
})
counter.update((v) => v + 1)
counter.unsubscribe(item1)
counter.update((v) => v + 2)
