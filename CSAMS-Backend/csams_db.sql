-- MySQL dump 10.13  Distrib 5.7.26, for Win64 (x86_64)
--
-- Host: localhost    Database: csams_db
-- ------------------------------------------------------
-- Server version	5.7.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `activity_association_models`
--

DROP TABLE IF EXISTS `activity_association_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activity_association_models` (
  `activity_id` bigint(20) unsigned DEFAULT NULL,
  `association_id` bigint(20) unsigned DEFAULT NULL,
  KEY `fk_activity_association_models_association_model` (`association_id`),
  KEY `fk_activity_association_models_activity_model` (`activity_id`),
  CONSTRAINT `fk_activity_association_models_activity_model` FOREIGN KEY (`activity_id`) REFERENCES `activity_models` (`id`),
  CONSTRAINT `fk_activity_association_models_association` FOREIGN KEY (`association_id`) REFERENCES `association_models` (`id`),
  CONSTRAINT `fk_activity_association_models_association_model` FOREIGN KEY (`association_id`) REFERENCES `association_models` (`id`),
  CONSTRAINT `fk_activity_models_activity_association` FOREIGN KEY (`activity_id`) REFERENCES `activity_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity_association_models`
--

LOCK TABLES `activity_association_models` WRITE;
/*!40000 ALTER TABLE `activity_association_models` DISABLE KEYS */;
INSERT INTO `activity_association_models` VALUES (123,122),(124,122),(125,122),(127,123),(128,123),(129,123),(130,124),(131,124),(132,124),(133,125),(134,125),(135,125),(136,126),(137,126),(138,126),(139,127),(140,127),(141,127),(142,128),(143,128),(144,128),(145,129),(146,129),(147,129),(148,130),(149,130),(150,130),(151,131),(152,131);
/*!40000 ALTER TABLE `activity_association_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `activity_log_models`
--

DROP TABLE IF EXISTS `activity_log_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activity_log_models` (
  `activity_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  KEY `fk_activity_log_models_activity_model` (`activity_id`),
  KEY `fk_activity_log_models_user` (`user_id`),
  CONSTRAINT `fk_activity_log_models_activity_model` FOREIGN KEY (`activity_id`) REFERENCES `activity_models` (`id`),
  CONSTRAINT `fk_activity_log_models_user` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_activity_log_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`),
  CONSTRAINT `fk_activity_models_activity_log` FOREIGN KEY (`activity_id`) REFERENCES `activity_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity_log_models`
--

LOCK TABLES `activity_log_models` WRITE;
/*!40000 ALTER TABLE `activity_log_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `activity_log_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `activity_models`
--

DROP TABLE IF EXISTS `activity_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activity_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `activity_name` varchar(32) COLLATE utf8_unicode_ci DEFAULT NULL,
  `start_time` datetime(3) DEFAULT NULL,
  `end_time` datetime(3) DEFAULT NULL,
  `location` varchar(32) COLLATE utf8_unicode_ci DEFAULT NULL,
  `introduction` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `image` longtext COLLATE utf8_unicode_ci,
  `responsible_person` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `tel` bigint(20) unsigned DEFAULT NULL,
  `score` bigint(20) unsigned DEFAULT NULL,
  `limit` json DEFAULT NULL,
  `is_end` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activity_models`
--

LOCK TABLES `activity_models` WRITE;
/*!40000 ALTER TABLE `activity_models` DISABLE KEYS */;
INSERT INTO `activity_models` VALUES (123,'健康讲座','2023-06-06 08:00:00.000','2023-06-17 08:00:00.000','学生活动中心','本次讲座将邀请专业医生为大家分享健康养生知识，欢迎踊跃参加！','/uploads/file/01dcb15c6f5f5fa801213f267b4a8d_20240607232852.jpg','李明',18812345678,3,'[\"其他\"]',0),(124,'摄影比赛','2022-07-12 08:00:00.000','2023-07-15 08:00:00.000','校园内','展现你的摄影才华，参加我们的摄影比赛，有机会获得丰厚奖品哦！','/uploads/file/true_20240607233045.jpg','李明',18812345678,4,'[\"其他\"]',0),(125,'志愿者招募','2022-08-20 08:00:00.000','2023-08-25 08:00:00.000','各社区','招募志愿者参与社区服务活动，一起为社会贡献爱心与力量。','/uploads/file/01f90a60bb232f11013eaf70e99e57_20240607233354.jpg','李明',18812345678,5,'[\"其他\"]',0),(127,'科技创新大赛','2022-10-05 08:00:00.000','2023-10-08 08:00:00.000','实验楼','展现你的科技创新成果，参加我们的科技创新大赛，赢取科技大奖！','/uploads/file/55b0132e38cbdcd0d199331ce4716874_20240607233528.jpg','王强',18823456789,5,'[\"其他\", \"人工智能\"]',0),(128,'环保公益活动','2022-11-15 08:00:00.000','2023-11-18 08:00:00.000','校园周边','加入我们的环保公益活动，为美丽的校园尽一份环保力量！','/uploads/file/v2-243d3e0d8c817ed46b760d73682dee55_r_20240607233609.jpg','王强',18823456789,3,'[\"环境工程\"]',0),(129,'文化交流沙龙','2022-12-03 08:00:00.000','2023-12-04 08:00:00.000','国际交流中心','来参加我们的文化交流沙龙，了解世界各地的文化魅力与风情！','/uploads/file/0146965d186afaa8012051cdc04c0f_20240607233703.jpg','王强',18823456789,4,'[\"其他\"]',0),(130,'心理健康讲座','2022-01-20 08:00:00.000','2023-01-21 08:00:00.000','心理健康中心','关注心理健康，参加我们的心理健康讲座，学习关爱自己与他人。','/uploads/file/a16fdffa9dcd82a6328a89cb9501d2c5_20240607233854.jpg','张伟',18834567890,5,'[\"其他\"]',0),(131,'创业论坛','2022-02-15 08:00:00.000','2023-02-16 08:00:00.000','创业孵化中心','踏入创业的世界，参加我们的创业论坛，与成功创业者面对面交流。','/uploads/file/a16fdffa9dcd82a6328a89cb9501d2c5_20240607234103.jpg','张伟',18834567890,4,'[\"其他\"]',0),(132,'读书分享会','2023-03-10 08:00:00.000','2025-03-11 08:00:00.000','图书馆报告厅','与书为伴，参加我们的读书分享会，分享阅读的快乐与感悟。','/uploads/file/true_20240607234112.jpg','张伟',18834567890,3,'[\"其他\", \"物联网工程\"]',0),(133,'科学实验竞赛','2023-04-18 08:00:00.000','2025-04-20 08:00:00.000','实验室','探索科学的奥秘，参加我们的科学实验竞赛，开启科学之旅！','/uploads/file/v2-2b53c49dad3f4b61ec96b5dded3b43b3_r_20240607234259.jpg','刘涛',18845678901,5,'[\"其他\", \"软件工程\"]',0),(134,'公益义卖','2023-05-22 08:00:00.000','2025-05-25 08:00:00.000','校园广场','参与我们的公益义卖，用爱心助力需要帮助的人群。','/uploads/file/v2-6a7c623f42d3c99334e6eb10aa2c6c2f_r_20240607234505.jpg','刘涛',18845678901,4,'[\"其他\"]',0),(135,'艺术展览','2023-06-18 08:00:00.000','2025-06-20 08:00:00.000','艺术馆','欣赏艺术之美，参观我们的艺术展览，感受艺术的魅力与情感。','/uploads/file/aac3075e15c9077dcf2f92e47fa4ffd3_20240607234542.jpg','刘涛',18845678901,5,'[\"其他\"]',0),(136,'健身挑战赛','2023-07-14 08:00:00.000','2025-07-17 08:00:00.000','大学体育馆','挑战自我，参加我们的健身挑战赛，释放健康与活力。','/uploads/file/01dcb15c6f5f5fa801213f267b4a8d_20240607234652.jpg','赵军',18856789012,4,'[\"其他\"]',0),(137,'国际文化节','2023-08-20 08:00:00.000','2025-08-23 08:00:00.000','校园内','融入世界，参与我们的国际文化节，感受不同国家的风情与风俗。','/uploads/file/0146965d186afaa8012051cdc04c0f_20240607234849.jpg','赵军',18856789012,5,'[\"其他\", \"英语\"]',0),(138,'心灵抒发夜','2023-09-15 08:00:00.000','2025-09-16 08:00:00.000','学生活动中心','释放心灵，参加我们的心灵抒发夜，倾诉内心的情感与烦恼。','/uploads/file/01dcb15c6f5f5fa801213f267b4a8d_20240607235003.jpg','赵军',18856789012,4,'[\"其他\"]',0),(139,'创业训练营','2023-10-10 08:00:00.000','2025-10-13 08:00:00.000','创业孵化中心','启程创业之旅，参与我们的创业训练营，开启创业之路。','/uploads/file/55b0132e38cbdcd0d199331ce4716874_20240607235115.jpg','周杰',18867890123,5,'[\"其他\"]',0),(140,'诗歌朗诵会','2023-11-18 08:00:00.000','2025-11-20 08:00:00.000','文学社活动室','用诗歌述说心声，参与我们的诗歌朗诵会，感受诗情画意。','/uploads/file/01f90a60bb232f11013eaf70e99e57_20240607235151.jpg','周杰',18867890123,4,'[\"其他\"]',0),(141,'科学知识竞赛','2023-12-14 08:00:00.000','2025-12-17 08:00:00.000','校园内','检验科学智慧，参加我们的科学知识竞赛，挑战科学极限。','/uploads/file/aac3075e15c9077dcf2f92e47fa4ffd3_20240607235223.jpg','周杰',18867890123,5,'[\"其他\"]',0),(142,'慈善义跑','2025-01-18 08:00:00.000','2026-01-20 08:00:00.000','校园操场','奔跑公益，参与我们的慈善义跑，用行动传递爱与温暖。','/uploads/file/true_20240607235326.jpg','吴磊',18878901234,4,'[\"其他\"]',0),(143,'国际交流晚会','2025-02-22 08:00:00.000','2026-02-25 08:00:00.000','国际交流中心','欢迎参加我们的国际交流晚会，感受世界各国的多彩风情与文化。','/uploads/file/true_20240607235403.jpg','吴磊',18878901234,5,'[\"其他\"]',0),(144,'心理健康工作坊','2025-03-20 08:00:00.000','2026-03-21 08:00:00.000','心理健康中心','关注心理健康，参加我们的心理健康工作坊，学习心理调适的方法。','/uploads/file/v2-2b53c49dad3f4b61ec96b5dded3b43b3_r_20240607235437.jpg','吴磊',18878901234,4,'[\"其他\"]',0),(145,'创业路演','2025-04-14 08:00:00.000','2026-04-16 08:00:00.000','创业孵化中心','展示创业成果，参加我们的创业路演，与投资人面对面沟通。','/uploads/file/v2-243d3e0d8c817ed46b760d73682dee55_r_20240607235650.jpg','孙健',18890123456,5,'[\"其他\"]',0),(146,'文学创作比赛','2025-05-18 08:00:00.000','2026-05-20 08:00:00.000','图书馆报告厅','展现文学才华，参加我们的文学创作比赛，展示文字的力量。','/uploads/file/a16fdffa9dcd82a6328a89cb9501d2c5_20240607235725.jpg','孙健',18890123456,4,'[\"其他\"]',0),(147,'科普知识竞赛','2025-06-22 08:00:00.000','2026-06-25 08:00:00.000','校园内','探索科学奥秘，参加我们的科普知识竞赛，普及科学知识。','/uploads/file/01dcb15c6f5f5fa801213f267b4a8d_20240607235853.jpg','孙健',18890123456,5,'[\"其他\"]',0),(148,'公益服务周','2025-07-20 08:00:00.000','2026-07-23 08:00:00.000','社区','关爱社会，参加我们的公益服务周，为社区居民奉献爱心。','/uploads/file/true_20240608000019.jpg','钱伟',18801234567,4,'[\"其他\"]',0),(149,'舞蹈表演','2025-08-15 08:00:00.000','2026-08-17 08:00:00.000','大学体育馆','释放舞蹈热情，参加我们的舞蹈表演，展现舞台魅力。','/uploads/file/55b0132e38cbdcd0d199331ce4716874_20240608000054.jpg','钱伟',18801234567,5,'[\"其他\"]',0),(150,'国际艺术交流展','2025-09-10 08:00:00.000','2026-09-12 08:00:00.000','艺术馆','欢迎参加我们的国际艺术交流展，领略世界各地的艺术风采。','/uploads/file/0146965d186afaa8012051cdc04c0f_20240608000125.jpg','钱伟',18801234567,4,'[\"其他\"]',0),(151,'健康生活讲座','2025-10-05 08:00:00.000','2026-10-08 08:00:00.000','学生活动中心','关注健康生活方式，参加我们的健康生活讲座，学习保健知识。','/uploads/file/v2-2b53c49dad3f4b61ec96b5dded3b43b3_r_20240608000300.jpg','朱伟',18812345012,5,'[\"其他\"]',0),(152,'职业规划讲座','2025-06-10 08:00:00.000','2026-06-10 08:00:00.000','教学楼A101','本次讲座将邀请资深职业规划师为大家分享就业指导和职业规划建议，欢迎踊跃参加','/uploads/file/v2-2b53c49dad3f4b61ec96b5dded3b43b3_r_20240608000406.jpg','朱伟',18812345012,4,'[\"其他\"]',0);
/*!40000 ALTER TABLE `activity_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `assignment_models`
--

DROP TABLE IF EXISTS `assignment_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `assignment_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `content` longtext COLLATE utf8_unicode_ci,
  `activity_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `is_submit` tinyint(1) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  `is_correct` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_activity_models_assignments` (`activity_id`),
  KEY `fk_user_models_assignments` (`user_id`),
  CONSTRAINT `fk_activity_models_assignments` FOREIGN KEY (`activity_id`) REFERENCES `activity_models` (`id`),
  CONSTRAINT `fk_user_models_assignments` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assignment_models`
--

LOCK TABLES `assignment_models` WRITE;
/*!40000 ALTER TABLE `assignment_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `assignment_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `association_member_models`
--

DROP TABLE IF EXISTS `association_member_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `association_member_models` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `association_id` bigint(20) unsigned DEFAULT NULL,
  `posts` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `joining_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_models_association_member` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2022304021 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `association_member_models`
--

LOCK TABLES `association_member_models` WRITE;
/*!40000 ALTER TABLE `association_member_models` DISABLE KEYS */;
INSERT INTO `association_member_models` VALUES (10123,131,'负责老师','2024-06-07 16:31:00.570'),(11234,132,'负责老师','2024-06-07 16:31:33.449'),(12345,122,'负责老师','2024-06-07 16:24:23.856'),(12347,133,'负责老师','2024-06-07 16:31:56.692'),(23456,123,'负责老师','2024-06-07 16:25:03.901'),(34567,124,'负责老师','2024-06-07 16:27:29.705'),(45678,125,'负责老师','2024-06-07 16:28:01.060'),(56789,126,'负责老师','2024-06-07 16:28:30.774'),(67890,127,'负责老师','2024-06-07 16:28:58.738'),(78901,128,'负责老师','2024-06-07 16:29:26.863'),(89012,129,'负责老师','2024-06-07 16:29:58.117'),(90123,130,'负责老师','2024-06-07 16:30:35.473'),(2022304001,122,'成员','2024-06-07 16:40:00.099'),(2022304002,122,'成员','2024-06-07 16:40:03.855'),(2022304003,122,'成员','2024-06-07 16:40:06.404'),(2022304004,122,'成员','2024-06-07 16:40:08.898'),(2022304005,122,'成员','2024-06-07 16:40:11.919'),(2022304006,122,'成员','2024-06-07 16:40:15.248'),(2022304007,122,'成员','2024-06-07 16:40:17.245'),(2022304008,122,'成员','2024-06-07 16:40:18.858'),(2022304009,122,'成员','2024-06-07 16:40:20.246'),(2022304010,122,'成员','2024-06-07 16:40:28.564'),(2022304011,122,'成员','2024-06-07 16:40:31.700'),(2022304012,122,'成员','2024-06-07 16:40:33.510'),(2022304013,122,'成员','2024-06-07 16:40:35.396'),(2022304014,122,'成员','2024-06-07 16:40:36.818'),(2022304015,122,'成员','2024-06-07 16:40:39.127'),(2022304016,123,'成员','2024-06-07 16:46:51.913'),(2022304017,123,'成员','2024-06-07 16:46:55.011'),(2022304018,123,'成员','2024-06-07 16:46:58.433'),(2022304019,123,'成员','2024-06-07 16:47:01.124'),(2022304020,123,'成员','2024-06-07 16:47:04.350');
/*!40000 ALTER TABLE `association_member_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `association_models`
--

DROP TABLE IF EXISTS `association_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `association_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `association_name` varchar(32) COLLATE utf8_unicode_ci DEFAULT NULL,
  `create_at` datetime(3) DEFAULT NULL,
  `teacher_name` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `teacher_id` bigint(20) unsigned DEFAULT NULL,
  `president` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `introduction` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_association_models_teacher` (`teacher_id`),
  CONSTRAINT `fk_association_models_teacher` FOREIGN KEY (`teacher_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=134 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `association_models`
--

LOCK TABLES `association_models` WRITE;
/*!40000 ALTER TABLE `association_models` DISABLE KEYS */;
INSERT INTO `association_models` VALUES (122,'科技协会','2024-06-07 16:24:23.851','李明',12345,'','欢迎加入科技协会，我们致力于推动科技创新和知识分享。'),(123,'社团联合','2024-06-07 16:25:03.889','王强',23456,'','这是一个致力于促进校园文化交流和学术探讨的协会。'),(124,'科技创新协会','2024-06-07 16:27:29.700','张伟',34567,'','该协会致力于推动科技创新，鼓励学生参与科研项目和技术竞赛。'),(125,'艺术表演团队','2024-06-07 16:28:01.055','刘涛',45678,'','这个团队聚集了校内优秀的艺术爱好者，致力于举办各种艺术表演和文化活动。'),(126,'环保协会','2024-06-07 16:28:30.768','赵军',56789,'','该协会旨在提倡环保意识，组织各种环保活动，为打造绿色校园贡献力量。'),(127,'志愿者服务团','2024-06-07 16:28:58.727','周杰',67890,'','这个团队志在服务社会，组织各种志愿活动，帮助有需要的人群。'),(128,'体育健身协会','2024-06-07 16:29:26.858','吴磊',78901,'','该协会鼓励学生积极参与体育运动，组织各类健身活动，提升身心健康。'),(129,'国际交流协会','2024-06-07 16:29:58.108','孙健',89012,'','这个协会旨在促进国际间的文化交流与友谊，举办各种国际文化活动。'),(130,'心理健康协会','2024-06-07 16:30:35.468','钱伟',90123,'','该协会致力于关注学生心理健康，开展心理辅导和心理健康教育活动。'),(131,'创业发展协会','2024-06-07 16:31:00.564','朱伟',10123,'','这个协会为有创业梦想的学生提供创业指导和资源支持，助力他们实现创业梦想。'),(132,'文学社','2024-06-07 16:31:33.440','杨洋',11234,'','该社团聚焦文学领域，组织文学创作、读书分享等活动，促进学生文学素养提升。'),(133,'科普社','2024-06-07 16:31:56.687','肖华',12347,'','这个社团致力于科普宣传，开展科学实验、科普讲座等活动，普及科学知识。');
/*!40000 ALTER TABLE `association_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `login_data_models`
--

DROP TABLE IF EXISTS `login_data_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `login_data_models` (
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `ip` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(42) COLLATE utf8_unicode_ci DEFAULT NULL,
  `token` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `device` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `addr` varchar(64) COLLATE utf8_unicode_ci DEFAULT NULL,
  KEY `fk_login_data_models_user_model` (`user_id`),
  CONSTRAINT `fk_login_data_models_user_model` FOREIGN KEY (`user_id`) REFERENCES `user_models` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `login_data_models`
--

LOCK TABLES `login_data_models` WRITE;
/*!40000 ALTER TABLE `login_data_models` DISABLE KEYS */;
INSERT INTO `login_data_models` VALUES (12345,'127.0.0.1','李明','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjM0NSwibmFtZSI6IuadjuaYjiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDYwNi45ODg1OTEyLCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.K8_-Bw78XRe3K0k-h7xZ_b6l1x2QO3EJ5HjNGpqqdAA','','内网地址'),(23456,'127.0.0.1','王强','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMzQ1NiwibmFtZSI6IueOi-W8uiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDY4MC41MzQ0ODIsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.n1QKYgDaJZmHKEf_r11q4-UobTkzPiAu9LMMr8zDemE','','内网地址'),(34567,'127.0.0.1','张伟','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozNDU2NywibmFtZSI6IuW8oOS8nyIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDgxMy41OTg5NTkyLCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.qKJ0iswOYAW8TZKL1IHPSlyTE2lQlks2soKVZGHS4h4','','内网地址'),(45678,'127.0.0.1','刘涛','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0NTY3OCwibmFtZSI6IuWImOa2myIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDg2My4zNDA2MDI5LCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.0Zlyi006luc3bB1k61BqtuppnydwdkHMS0adrPkGblk','','内网地址'),(56789,'127.0.0.1','赵军','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1Njc4OSwibmFtZSI6Iui1teWGmyIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDg5NC44Njk4NzIsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.Xw0CW_ifhdq-9qLKBHD64N_TwSjeHkfz0CnMUlouu_c','','内网地址'),(67890,'127.0.0.1','周杰','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2Nzg5MCwibmFtZSI6IuWRqOadsCIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDkyMy4xMjcxNywiaXNzIjoic2NjXHUwMDI2enpyXHUwMDI2eHp4XHUwMDI2emp5XHUwMDI2ZnNjIn0.QHBkIGBeJNtvsaRTNc92jwdB6Sxgo78R9J-Ri1WUAKo','','内网地址'),(78901,'127.0.0.1','吴磊','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3ODkwMSwibmFtZSI6IuWQtOejiiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDk0OS4xOTkyNTg4LCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.kwTZ6aE5sieGNI0AyXc5GT1vtAUQC81AhmjZ7kg5_HI','','内网地址'),(89012,'127.0.0.1','孙健','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4OTAxMiwibmFtZSI6IuWtmeWBpSIsInJvbGUiOjEsImV4cCI6MTcxNzgyMDk4Mi4yNjQwNzYsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.fCd-IummKIR2pIqFaUIKCNVrSykEkNxmcVPYIflE8SQ','','内网地址'),(90123,'127.0.0.1','钱伟','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MDEyMywibmFtZSI6IumSseS8nyIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTAwOS40Mzc0NjEsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.QqnPYkNaZFxleu1sC8ugRKQj6uSPvXv52uuyUxoibe4','','内网地址'),(10123,'127.0.0.1','朱伟','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDEyMywibmFtZSI6IuacseS8nyIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTA0Ni43MDU2OTMsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.eH_wbO_rQzs1u_I_axGtR6cKsi9P53G836v7Bts2aKQ','','内网地址'),(11234,'127.0.0.1','杨洋','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMTIzNCwibmFtZSI6IuadqOa0iyIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTA3Ny41OTY0OTU5LCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.ldNloRVrb2Jv-qGgXgK4fYMn-zEqjKFhSkh2eCjZ3I0','','内网地址'),(12347,'127.0.0.1','肖华','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjM0NywibmFtZSI6IuiCluWNjiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTEwMy41NTA5NDUsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.hz4AH_MfY0GpNXxEPApZJdYwU8G_6f3dajWZrPZu-NY','','内网地址'),(12345,'127.0.0.1','李明','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjM0NSwibmFtZSI6IuadjuaYjiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTM0OS41NzU0OTUsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.GBG5Q0CFqRHdSO5t47-zM_hSq0dSBsDBFpoHxn3cIRw','','内网地址'),(23456,'127.0.0.1','王强','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMzQ1NiwibmFtZSI6IueOi-W8uiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTY2NS45NTE1NTMsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.iPn0xPt37l9VLi3Py82-sH0KCKWUau2XYggYKWyjMsU','','内网地址'),(12345,'127.0.0.1','李明','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjM0NSwibmFtZSI6IuadjuaYjiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTkwOC4wNDAwNjIsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.YvpdUQPTguQh-aZYU1spqhZewlcea5i0ZVc5IOo_gaM','','内网地址'),(23456,'127.0.0.1','王强','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMzQ1NiwibmFtZSI6IueOi-W8uiIsInJvbGUiOjEsImV4cCI6MTcxNzgyMTk4Ny4xNzQ3MzQsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.GmquIZMVoSf3iiJ44yJkC8yQ7UkecW9giY__eGTtOWk','','内网地址'),(12345,'127.0.0.1','李明','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjM0NSwibmFtZSI6IuadjuaYjiIsInJvbGUiOjEsImV4cCI6MTcxNzg0NTc2Mi44MDY3NjQxLCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.Cej-5iZb-8ZYSs87VE3MDTSq4ghe4YvfOaD69gn1T-Q','','内网地址'),(23456,'127.0.0.1','王强','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMzQ1NiwibmFtZSI6IueOi-W8uiIsInJvbGUiOjEsImV4cCI6MTcxNzg0NjQ3MC4yMzcxNjgsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.GJnUtyjIi5zIRBZFmQBie9ZIyziJHGeSVLFAXEJXynA','','内网地址'),(34567,'127.0.0.1','张伟','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozNDU2NywibmFtZSI6IuW8oOS8nyIsInJvbGUiOjEsImV4cCI6MTcxNzg0NjY5OC42OTIwMSwiaXNzIjoic2NjXHUwMDI2enpyXHUwMDI2eHp4XHUwMDI2emp5XHUwMDI2ZnNjIn0.RtxUD4X7Po7F7vZbh-mgSqIwsNHeA0odbJrt3qPTajk','','内网地址'),(45678,'127.0.0.1','刘涛','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0NTY3OCwibmFtZSI6IuWImOa2myIsInJvbGUiOjEsImV4cCI6MTcxNzg0NjkzNi4wMjk5ODYxLCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.Ii0lOzhxWJMaL_61fd-h0_llG0CGyoruU2vnqo98DeQ','','内网地址'),(56789,'127.0.0.1','赵军','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1Njc4OSwibmFtZSI6Iui1teWGmyIsInJvbGUiOjEsImV4cCI6MTcxNzg0NzE2Mi4zODgzODMsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.E1rIM9XsOWKdfn9IJILkx7FEusg7TSXyIg2lqjp8NEw','','内网地址'),(67890,'127.0.0.1','周杰','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2Nzg5MCwibmFtZSI6IuWRqOadsCIsInJvbGUiOjEsImV4cCI6MTcxNzg0NzQ0Mi40MTE1NTQsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.bbk-sItn4N9mlPFiLCmlMkZ4f7_yqxG2_gUrg-RuTC4','','内网地址'),(78901,'127.0.0.1','吴磊','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3ODkwMSwibmFtZSI6IuWQtOejiiIsInJvbGUiOjEsImV4cCI6MTcxNzg0NzU3Mi45NDUzNTE4LCJpc3MiOiJzY2NcdTAwMjZ6enJcdTAwMjZ4enhcdTAwMjZ6anlcdTAwMjZmc2MifQ.f0sIQOqiHzt8GaQTGKIMH7MMvoVVkppzmmV_DBh-S-Q','','内网地址'),(89012,'127.0.0.1','孙健','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4OTAxMiwibmFtZSI6IuWtmeWBpSIsInJvbGUiOjEsImV4cCI6MTcxNzg0NzY5NS40NzcxMDMsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.z-7egIt_drzgTQoubrYlPxW4hVMERWgkfe5ubo2wpPE','','内网地址'),(90123,'127.0.0.1','钱伟','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MDEyMywibmFtZSI6IumSseS8nyIsInJvbGUiOjEsImV4cCI6MTcxNzg0Nzk1NS4yNDcwMywiaXNzIjoic2NjXHUwMDI2enpyXHUwMDI2eHp4XHUwMDI2emp5XHUwMDI2ZnNjIn0.aoB-XGy93MtOI4T9OR8ta6xYLaSjaXX_Gbi0IfLvujY','','内网地址'),(10123,'127.0.0.1','朱伟','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDEyMywibmFtZSI6IuacseS8nyIsInJvbGUiOjEsImV4cCI6MTcxNzg0ODExMC4zNTkyOTcsImlzcyI6InNjY1x1MDAyNnp6clx1MDAyNnh6eFx1MDAyNnpqeVx1MDAyNmZzYyJ9.jkXHQa68uTWPD09_BEZt35LwXA6KcB2yWaIDyiD5bdw','','内网地址');
/*!40000 ALTER TABLE `login_data_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_models`
--

DROP TABLE IF EXISTS `user_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_models` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `password` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `age` tinyint(4) DEFAULT NULL,
  `gender` varchar(3) COLLATE utf8_unicode_ci DEFAULT NULL,
  `avatar` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `role` tinyint(4) DEFAULT NULL,
  `major` tinyint(4) DEFAULT NULL,
  `tel` bigint(20) unsigned DEFAULT NULL,
  `integrity` bigint(20) DEFAULT NULL,
  `score` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2022304021 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_models`
--

LOCK TABLES `user_models` WRITE;
/*!40000 ALTER TABLE `user_models` DISABLE KEYS */;
INSERT INTO `user_models` VALUES (10123,'$2a$04$NlA66wgdQqwV4dFa.0HHR.EYPmkVNqez2y5BEreYnly4rHHNORJxe','朱伟',34,'男','/uploads/avatar/default.png',1,40,18812345012,100,0),(11234,'$2a$04$FsBWBzf7duv2EpnHPjbm4OIlz0dxhW8tYOO7F56SQGt4fCpHmK.za','杨洋',35,'男','/uploads/avatar/default.png',1,44,18823450123,100,0),(12345,'$2a$04$vcQ7A.DtuUKhy5hmNuLL9.L8.dQ0CHJ0v3kySc9EWwqYKmdhsXdMW','李明',25,'男','/uploads/avatar/default.png',1,60,18812345678,100,0),(12347,'$2a$04$jm3R8H6Zf0k.9vDUxmSy6eJJZVKsgSdRZpdkVPfg/a9PdUq9J5aIW','肖华',36,'男','/uploads/avatar/default.png',1,49,18834501234,100,0),(23456,'$2a$04$Jf6gzmX9Y7PV5jL1r9wKouia7WxT8C/wNUXY.jDxHnAF9DehuPjv.','王强',28,'男','/uploads/avatar/default.png',1,21,18823456789,100,0),(34567,'$2a$04$yY3u1zL0E6DeLoqn7H.NJ.oBAhBVrRzD9CGHCMwr75GrWy7GVe46y','张伟',30,'男','/uploads/avatar/default.png',1,35,18834567890,100,0),(45678,'$2a$04$7s1U95w0/uDdJuUwIg32h.CUQDIw5lM2YPSbQfJNN2alWso2oLkK2','刘涛',26,'男','/uploads/avatar/default.png',1,38,18845678901,100,0),(56789,'$2a$04$M8tTi5GdaJgJapxoIViFVel6FdR/OCWp0iCzji5o9jvOHQm2BHox6','赵军',27,'男','/uploads/avatar/default.png',1,48,18856789012,100,0),(67890,'$2a$04$bCyfuFwdb2j0/LGOzP9H7OBq/o0X0sUU3tH2iGqqphYPunjkgnEpG','周杰',29,'男','/uploads/avatar/default.png',1,20,18867890123,100,0),(78901,'$2a$04$s4iE2zYwlELMi/zS22K3Kee9pt/Jt/9rMjx0GuZzhbQJganBvxthK','吴磊',31,'男','/uploads/avatar/default.png',1,19,18878901234,100,0),(89012,'$2a$04$lekx97k2o4Roke8L/kIPh.tHNTGoDz4s7pQ67oNo.ylAbgkhJq3uq','孙健',32,'男','/uploads/avatar/default.png',1,14,18890123456,100,0),(90123,'$2a$04$14xaXWDtdLGE8BGyvoeo2.rD0GOOn5gU8YsydGOBfPUFwQndVacFm','钱伟',33,'男','/uploads/avatar/default.png',1,28,18801234567,100,0),(2022304001,'$2a$04$aQK04J1hZq.ye3mNd3I7JOPlZmZq.6x7PkT0RL2qNa.O6AUYLkALG','张三',20,'男','/uploads/avatar/default.png',2,59,18812345679,100,0),(2022304002,'$2a$04$9riQc0xENTOY0sJxAKotMOWyodQ2T5gYf4dsGZ9qOk87m.3vsFO46','李四',21,'男','/uploads/avatar/default.png',2,10,18823456780,100,0),(2022304003,'$2a$04$c66CHwALccn7aOKNq1LHkOjb9Pz6pMhG/.wxVrbptlDcD15rG.hBy','王五',22,'男','/uploads/avatar/default.png',2,51,18834567891,100,0),(2022304004,'$2a$04$3G2nCDr5NYkppsTFnfbwku2Te.NyddjgHUc3zXGR14TyYn26sjVPe','赵六',23,'男','/uploads/avatar/default.png',2,62,18845678902,100,0),(2022304005,'$2a$04$0pOYwu5S5dRzXwwZ/TXli.2X3MC9TBgW.lk9dZgqln7ZpWmK6hrMS','孙七',24,'男','/uploads/avatar/default.png',2,28,18856789013,100,0),(2022304006,'$2a$04$iBHjFEtZDq/QOC9S7SRtB.37YV5GEDOXLp8GMqE55pr5DzGH5iGwK','周八',25,'男','/uploads/avatar/default.png',2,34,18867890124,100,0),(2022304007,'$2a$04$iX6vAxa7fFyemxa6GM.xbuyuStk.CHhbs3r11W/FcdkpH/bxfab8i','吴九',26,'男','/uploads/avatar/default.png',2,39,18878901235,100,0),(2022304008,'$2a$04$/br2GnB87Sjr5Pg7z2Ugvuvgz8vqP6m/yHuSnVFvE6k5UMIYcL5Yq','郑十',27,'男','/uploads/avatar/default.png',2,35,18889012346,100,0),(2022304009,'$2a$04$990BbE8LenAKWZ0YYjmTvO.7H7Xfhe2KE5InMnEPle8JAnVagIwke','刘十一',28,'男','/uploads/avatar/default.png',2,10,18890123457,100,0),(2022304010,'$2a$04$f1G65wUiz7IjiXd7C3iCIOLZK4zcZiCEKC/6gmWhb/gYzMw3b4LlW','朱十二',29,'男','/uploads/avatar/default.png',2,6,18801234568,100,0),(2022304011,'$2a$04$64VOwaqaNAcESp2oci0sr.fa2gJOMOpxRMjG27mP4kDS.M2xGhHBm','李十三',30,'男','/uploads/avatar/default.png',2,12,18812345679,100,0),(2022304012,'$2a$04$EgfwY6B7Hz11eRirIebLY.2ZDw/ETz1DusgT.4LGkn9H9kNV4RFii','王十四',31,'男','/uploads/avatar/default.png',2,21,18823456780,100,0),(2022304013,'$2a$04$wnjKW0a1Q83P4LaUOFYphOQMxkU3cZDD2qhjAVLtvLxaWNzjnKMTy','赵十五',32,'男','/uploads/avatar/default.png',2,49,18834567891,100,0),(2022304014,'$2a$04$/4iS5KX5MwbM0ZWuJgLJ.em9OKVZkivfNu6TXr1nUM7wtnwHAyjQq','孙十六',33,'男','/uploads/avatar/default.png',2,31,18845678902,100,0),(2022304015,'$2a$04$2q0.5eF/mOy2Vn5jxgXnMernR7ISSttoMvGninNEI8zpNPCadQ7L.','周十七',34,'男','/uploads/avatar/default.png',2,59,18856789013,100,0),(2022304016,'$2a$04$avMCAc/2hQ9hMjj7JUOxuONUPhbF3dizEdJs1nB97E8srAMNoTM/W','吴十八',35,'男','/uploads/avatar/default.png',2,48,18867890124,100,0),(2022304017,'$2a$04$s2yYBMdYhj2rutfn/ostJurREd8/R3rNMUlSll2IVuzHqZXX6fIom','郑十九',36,'男','/uploads/avatar/default.png',2,34,18878901235,100,0),(2022304018,'$2a$04$n8fm1Nid4H0pqIUFu32TRepTn1AKigAII0vXxaEFFEFbV1oZlKRA6','刘二十',37,'男','/uploads/avatar/default.png',2,45,18889012346,100,0),(2022304019,'$2a$04$h7X0987KfoWj0y8yQfAu2.cgnSU0cMO28H16X8xSC.sDp7MfV2yx2','朱二十一',38,'男','/uploads/avatar/default.png',2,10,18890123457,100,0),(2022304020,'$2a$04$Is52HvYP42v84ug6eZXd5uFlofDW.o3dpVzdzB7VBLogkDE4IBEVa','李二十二',39,'男','/uploads/avatar/default.png',2,16,18801234568,100,0);
/*!40000 ALTER TABLE `user_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-08  0:06:17
