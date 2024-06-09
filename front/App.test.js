import { render } from "@testing-library/svelte";
import App from "./App.svelte";

test("renders the component", () => {
  const { getByText } = render(App);
  expect(getByText("Craft&Combat")).toBeInTheDocument(); 
});