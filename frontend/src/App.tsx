import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import {
  ChakraProvider,
  Box,
  useColorModeValue,
  Drawer,
  DrawerContent,
} from "@chakra-ui/react";
import Home from "./routes/Home";
import AccessHistory from "./routes/AccessHistory";
import ChangePassword from "./routes/ChangePassword";
import Profile from "./routes/Profile";
import SignInWebauthn from "./routes/SignInWebauthn";
import RegisterWebauthn from "./routes/RegisterWebauthn";
import SignInDigest from "./routes/SignInDigest";
import SignUp from "./routes/SignUp";
import LabAssistant from "./routes/LabAssistant";
import Footer from "./features/Footer";
import SidebarContent from "./features/SidebarContent";
import MobileNav from "./features/MobileNav";
import { useDisclosure } from "@chakra-ui/react";
import { Ranking } from "./routes/Ranking";
import { UserProvider } from "./userContext"; // Import UserProvider
import { UserSetting } from "./routes/UserSetting";

const App = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <ChakraProvider>
      <UserProvider>
        <BrowserRouter>
          <Box
            minH="100vh"
            bg={useColorModeValue("gray.100", "gray.900")}
            overflowX={"hidden"}
          >
            <MobileNav onOpen={onOpen} />
            <SidebarContent
              onClose={onClose}
              display={{ base: "none", md: "block" }}
              pt={20}
            />
            <Drawer
              isOpen={isOpen}
              placement="left"
              onClose={onClose}
              returnFocusOnClose={false}
              onOverlayClick={onClose}
              size="full"
            >
              <DrawerContent>
                <SidebarContent onClose={onClose} />
              </DrawerContent>
            </Drawer>
            <Box
              pl={{ base: 2, md: 64 }}
              pr={{ base: 2, md: 6 }}
              pt={{ base: 2, md: 24 }}
            >
              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/access-history" element={<AccessHistory />} />
                <Route path="/profile" element={<Profile />} />
                <Route path="/ranking" element={<Ranking />} />
                <Route path="/sign-in-webauthn" element={<SignInWebauthn />} />
                <Route
                  path="/register-webauthn"
                  element={<RegisterWebauthn />}
                />
                <Route path="/lab-assistant" element={<LabAssistant />} />
                <Route path="/sign-in-digest" element={<SignInDigest />} />
                <Route path="/user-setting" element={<UserSetting />} />
                <Route path="/sign-up" element={<SignUp />} />
                <Route path="/change-password" element={<ChangePassword />} />
              </Routes>
            </Box>
          </Box>
          <Footer />
        </BrowserRouter>
      </UserProvider>
    </ChakraProvider>
  );
};

export default App;
