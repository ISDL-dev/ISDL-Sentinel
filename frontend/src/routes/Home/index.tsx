import { Button, Grid } from "@chakra-ui/react";
import "./Home.css";

function Home() {
  return (
    <Grid
      templateColumns="repeat(3, 1fr)"
      alignItems={"center"}
      w={"-moz-max-content"}
    >
      <h1 className="block mb-1 text-4xl font-bold text-gray-900 dark:text-white p-3 text-left">
        出席者一覧
      </h1>
      <Grid
        templateColumns="repeat(2, 1fr)"
        alignItems={"center"}
        w={"-moz-max-content"}
        column={3}
      >
        <Button colorScheme="teal" variant="solid" size="lg" width={36}>
          入室
        </Button>
        <Button colorScheme="teal" variant="solid" size="lg" width={36}>
          退室
        </Button>
      </Grid>
    </Grid>
  );
}

export default Home;
