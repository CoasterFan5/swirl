export { };

declare global {
  const swirl: {
    state<T>(value: T): State<T>;
  };


}
declare module "swirl" {
  const swirl: {
    state<T>(value: T): StateManager<T>;
  };
  export default swirl;
  export { StateManager };
}
