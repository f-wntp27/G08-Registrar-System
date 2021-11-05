import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบลงทะเบียนรายวิชา</h1>
        <h4><u>Requirements</u></h4>
        <p>
          ระบบลงทะเบียนเรียนของมหาวิทยาลัย 
          เป็นระบบที่ให้นักศึกษาสามารถ login เข้าระบบเพื่อทำการลงทะเบียนรายวิชาที่จะเรียนในแต่ละภาคการศึกษา 
          นักศึกษาสามารถลงทะเบียนรายวิชาต่าง ๆ โดยจะมีข้อมูลของรายวิชา 
          ได้แก่ ชื่อรายวิชา กลุ่มของรายวิชา อาจารย์ผู้สอน และหมายเหตุของรายวิชา 
          เมื่อนักศึกษาเลือกรายวิชาและระบุประเภทการลงทะเบียนได้เรียบร้อยแล้ว 
          สามารถกดบันทึกข้อมูลเพื่อแสดงรายวิชาที่นักศึกษาได้ลงทะเบียนไปแล้วในภาคการศึกษานั้น ๆ 
          นอกจากนี้นักศึกษาสามารถตรวจสอบประวัติการลงทะเบียนได้ที่ ใบบันทึกผลการลงทะเบียนเรียน 
          โดยจะแสดงรายวิชาที่เคยลงทะเบียนไปในแต่ละภาคการศึกษาได้
        </p>
        <br />
        <h4><u>User Story</u> (ระบบลงทะเบียนรายวิชา *)</h4>
        <p> 
          <b>ในบทบาทของ</b>	นักศึกษา<br />
          <b>ฉันต้องการ</b>	ให้ระบบสามารถบันทึกการลงทะเบียนของนักศึกษาในแต่ละภาคการศึกษา<br />
          <b>เพื่อ</b>	ให้ฉันสามารถบันทึกข้อมูลการลงทะเบียนรายวิชาไว้ในใบลงทะเบียนเรียน<br />
        </p>
      </Container>
    </div>
  );
}
export default Home;
