import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title = {'ยินดีตอนรับเข้าสู่ระบบประวัติแพทย์'}
      //  subtitle=""
     ></Header>
     <Content>
       <ContentHeader title="ข้อมูลแพทย์">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             บันทึกข้อมูลส่วนตัวแพทย์
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;