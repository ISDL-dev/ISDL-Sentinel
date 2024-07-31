import {
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  useColorModeValue,
} from "@chakra-ui/react";
import { useState } from "react";
import { RankingList } from "../../features/RankingList";

export const Ranking = () => {
  const colors = useColorModeValue(
    ["red.50", "blue.50"],
    ["red.900", "blue.900"]
  );
  const [tabIndex, setTabIndex] = useState(0);
  const bg = colors[tabIndex];
  return (
    <Tabs
      isFitted
      onChange={(index) => setTabIndex(index)}
      bg={bg}
      border={"2px"}
      borderColor="gray.200"
      mt={{ base: 20, md: 0 }}
      width="100%"
    >
      <TabList mb="1em">
        <Tab>滞在時間</Tab>
        <Tab>出席日数</Tab>
      </TabList>
      <TabPanels>
        <TabPanel key={"stay_time"}>
          <RankingList placeholder="stay_time"></RankingList>
        </TabPanel>
        <TabPanel key={"attendance_days"}>
          <RankingList placeholder="attendance_days"></RankingList>
        </TabPanel>
      </TabPanels>
    </Tabs>
  );
};
