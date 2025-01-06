CREATE TABLE "account" (
	"id" text PRIMARY KEY NOT NULL,
	"account_id" text NOT NULL,
	"provider_id" text NOT NULL,
	"user_id" text NOT NULL,
	"access_token" text,
	"refresh_token" text,
	"id_token" text,
	"expires_at" timestamp with time zone,
	"password" text
);
--> statement-breakpoint
CREATE TABLE "activity_types" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"name" text NOT NULL,
	"published" boolean DEFAULT false NOT NULL,
	"category" text,
	"card_color" text
);
--> statement-breakpoint
CREATE TABLE "assets" (
	"id" uuid PRIMARY KEY NOT NULL,
	"file_name" text NOT NULL,
	"content_type" text NOT NULL,
	"e_tag" text,
	"container_name" text NOT NULL,
	"uri" text NOT NULL,
	"size" integer NOT NULL,
	"metadata" jsonb,
	"is_public" boolean DEFAULT true NOT NULL,
	"published" boolean DEFAULT true NOT NULL,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	CONSTRAINT "assets_fileName_unique" UNIQUE("file_name")
);
--> statement-breakpoint
CREATE TABLE "chapters" (
	"chapter_id" uuid PRIMARY KEY NOT NULL,
	"lesson_id" uuid NOT NULL,
	"nav_item_name" text NOT NULL,
	"description" text,
	"order" integer,
	"content" jsonb,
	"published" boolean DEFAULT false NOT NULL,
	"title" text NOT NULL,
	"chapter_number" integer,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "classes_activities" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"term" text,
	"start_date" timestamp with time zone,
	"end_date" timestamp with time zone,
	"homework_id" uuid NOT NULL,
	"day" text,
	"period" text,
	"class_id" uuid NOT NULL,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "class_scores" (
	"class_id" uuid NOT NULL,
	"user_id" text NOT NULL,
	"score" integer NOT NULL,
	"week" integer NOT NULL,
	CONSTRAINT "class_scores_pkey" PRIMARY KEY("class_id","user_id","week")
);
--> statement-breakpoint
CREATE TABLE "class_users" (
	"class_id" uuid NOT NULL,
	"user_id" text NOT NULL
);
--> statement-breakpoint
CREATE TABLE "classes" (
	"class_id" uuid PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"description" text,
	"course_code" text NOT NULL,
	"duration" integer,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "courses" (
	"courses_id" uuid PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"description" text,
	"chapters" integer,
	"sections" integer,
	"duration" integer,
	"category" text,
	"image" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "educational_standards" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"code" text NOT NULL,
	"description" text NOT NULL,
	"subject" text NOT NULL,
	"category" text NOT NULL,
	"sub_category" text NOT NULL,
	"level" integer NOT NULL,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "educations_standards_questions_mapping" (
	"question_id" uuid NOT NULL,
	"standard_id" uuid NOT NULL,
	CONSTRAINT "educations_standards_questions_mapping_pkey" PRIMARY KEY("question_id","standard_id")
);
--> statement-breakpoint
CREATE TABLE "enrolments" (
	"organisation_id" uuid,
	"course_id" uuid,
	"user_id" text,
	"name" text NOT NULL,
	"tags" text[],
	"content" jsonb,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "groups_users" (
	"group_id" uuid,
	"user_id" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "groups" (
	"group_id" uuid PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"course_id" uuid,
	"tags" text[],
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "homework_questions" (
	"id" uuid PRIMARY KEY NOT NULL,
	"question_id" uuid,
	"homework_id" uuid
);
--> statement-breakpoint
CREATE TABLE "homework_submissions" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"user_id" text,
	"homework_id" uuid,
	"progress" integer DEFAULT 0,
	"content" jsonb,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "homework" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"name" text NOT NULL,
	"created_by" text,
	"tags" text[],
	"content" jsonb,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "users" (
	"id" text PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
	"fullname" text NOT NULL,
	"email" text NOT NULL,
	"email_verified" boolean DEFAULT false,
	"image" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	CONSTRAINT "users_email_unique" UNIQUE("email")
);
--> statement-breakpoint
CREATE TABLE "invitees" (
	"id" uuid PRIMARY KEY NOT NULL,
	"email" text NOT NULL,
	"name" text,
	"invited_by" text NOT NULL,
	"relation" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "jwks" (
	"id" text PRIMARY KEY NOT NULL,
	"public_key" text,
	"private_key" text,
	"created_at" timestamp DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "lesson_resources" (
	"lesson_id" uuid NOT NULL,
	"resource_id" uuid NOT NULL,
	CONSTRAINT "lesson_resources_pkey" PRIMARY KEY("lesson_id","resource_id")
);
--> statement-breakpoint
CREATE TABLE "lessons" (
	"lesson_id" uuid PRIMARY KEY NOT NULL,
	"section_id" uuid NOT NULL,
	"course_id" uuid,
	"title" text NOT NULL,
	"summary" text,
	"content" jsonb,
	"duration" integer,
	"order_in_subject" integer,
	"image" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	CONSTRAINT "unique_lesson_key" UNIQUE("title","section_id","course_id")
);
--> statement-breakpoint
CREATE TABLE "modules" (
	"id" bigint PRIMARY KEY NOT NULL,
	"title" text,
	"description" text,
	"asset" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "notes" (
	"note_id" uuid PRIMARY KEY NOT NULL,
	"user_id" text NOT NULL,
	"content" text,
	"lesson_id" uuid,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "organisations" (
	"organisations_id" uuid PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"address" text,
	"contact_details" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "organisation_users" (
	"organisation_id" uuid,
	"user_id" text,
	"role_id" integer,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "outcomes" (
	"outcome_id" serial PRIMARY KEY NOT NULL,
	"name" text NOT NULL,
	"code" text NOT NULL,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	CONSTRAINT "outcomes_code_key" UNIQUE("code")
);
--> statement-breakpoint
CREATE TABLE "profiles" (
	"id" uuid PRIMARY KEY NOT NULL,
	"username" text,
	"egl" integer,
	"enrolled_at" text,
	"terms_accepted" boolean DEFAULT false NOT NULL,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "questions" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"egl" integer,
	"curriculum_reference" text[],
	"cognitive_skill" text[],
	"reading_ability" integer,
	"writing_ability" integer,
	"listening_ability" integer,
	"title" text,
	"question" text,
	"answer" text,
	"options" jsonb,
	"type" text
);
--> statement-breakpoint
CREATE TABLE "quizzes" (
	"title" text,
	"description" text,
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"total_questions" integer,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "quizzes_questions" (
	"id" uuid PRIMARY KEY NOT NULL,
	"quizzes_id" uuid,
	"question" text,
	"answer" text,
	"options" jsonb,
	"type" text,
	"question_order" integer,
	"resources" jsonb[],
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "random_questions" (
	"id" uuid,
	"egl" integer,
	"curriculum_reference" text[],
	"cognitive_skill" text[],
	"reading_ability" integer,
	"writing_ability" integer,
	"listening_ability" integer,
	"title" text,
	"question" text,
	"answer" text,
	"options" jsonb,
	"type" text
);
--> statement-breakpoint
CREATE TABLE "resources" (
	"resource_id" uuid PRIMARY KEY NOT NULL,
	"resource_type" text,
	"size" bigint NOT NULL,
	"metadata" jsonb,
	"etag" text,
	"provider" text DEFAULT 'internal' NOT NULL,
	"uri" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "roles" (
	"role_id" serial PRIMARY KEY NOT NULL,
	"name" text,
	"description" text,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "sections" (
	"section_id" uuid PRIMARY KEY NOT NULL,
	"courses_id" uuid NOT NULL,
	"name" text NOT NULL,
	"description" text,
	"order" integer,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
CREATE TABLE "session" (
	"id" text PRIMARY KEY NOT NULL,
	"expires_at" timestamp NOT NULL,
	"ip_address" text,
	"user_agent" text,
	"user_id" text NOT NULL
);
--> statement-breakpoint
CREATE TABLE "user_roles" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"user_id" text NOT NULL,
	"role_id" serial NOT NULL,
	"context_id" uuid NOT NULL,
	"context_type" text NOT NULL,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL,
	CONSTRAINT "user_roles_user_id_role_key" UNIQUE("user_id","role_id")
);
--> statement-breakpoint
CREATE TABLE "verification" (
	"id" text PRIMARY KEY NOT NULL,
	"identifier" text NOT NULL,
	"value" text NOT NULL,
	"expires_at" timestamp NOT NULL,
	"created_at" timestamp
);
--> statement-breakpoint
CREATE TABLE "weekly_planner" (
	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
	"class_id" uuid NOT NULL,
	"term" text NOT NULL,
	"week" integer NOT NULL,
	"content" jsonb,
	"updated_at" timestamp with time zone DEFAULT now() NOT NULL,
	"created_at" timestamp with time zone DEFAULT now() NOT NULL
);
--> statement-breakpoint
ALTER TABLE "account" ADD CONSTRAINT "account_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "chapters" ADD CONSTRAINT "chapters_lesson_id_lessons_lesson_id_fk" FOREIGN KEY ("lesson_id") REFERENCES "public"."lessons"("lesson_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "classes_activities" ADD CONSTRAINT "classes_activities_homework_id_homework_id_fk" FOREIGN KEY ("homework_id") REFERENCES "public"."homework"("id") ON DELETE cascade ON UPDATE cascade;--> statement-breakpoint
ALTER TABLE "classes_activities" ADD CONSTRAINT "classes_activities_class_id_classes_class_id_fk" FOREIGN KEY ("class_id") REFERENCES "public"."classes"("class_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "class_scores" ADD CONSTRAINT "class_scores_class_id_classes_class_id_fk" FOREIGN KEY ("class_id") REFERENCES "public"."classes"("class_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "class_scores" ADD CONSTRAINT "class_scores_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "class_users" ADD CONSTRAINT "class_users_class_id_classes_class_id_fk" FOREIGN KEY ("class_id") REFERENCES "public"."classes"("class_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "class_users" ADD CONSTRAINT "class_users_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "educations_standards_questions_mapping" ADD CONSTRAINT "educations_standards_questions_mapping_question_id_questions_id_fk" FOREIGN KEY ("question_id") REFERENCES "public"."questions"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "educations_standards_questions_mapping" ADD CONSTRAINT "educations_standards_questions_mapping_standard_id_educational_standards_id_fk" FOREIGN KEY ("standard_id") REFERENCES "public"."educational_standards"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "enrolments" ADD CONSTRAINT "enrolments_organisation_id_organisations_organisations_id_fk" FOREIGN KEY ("organisation_id") REFERENCES "public"."organisations"("organisations_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "enrolments" ADD CONSTRAINT "enrolments_course_id_courses_courses_id_fk" FOREIGN KEY ("course_id") REFERENCES "public"."courses"("courses_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "enrolments" ADD CONSTRAINT "enrolments_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "groups_users" ADD CONSTRAINT "groups_users_group_id_groups_group_id_fk" FOREIGN KEY ("group_id") REFERENCES "public"."groups"("group_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "groups_users" ADD CONSTRAINT "groups_users_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "homework_questions" ADD CONSTRAINT "homework_questions_question_id_questions_id_fk" FOREIGN KEY ("question_id") REFERENCES "public"."questions"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "homework_questions" ADD CONSTRAINT "homework_questions_homework_id_homework_id_fk" FOREIGN KEY ("homework_id") REFERENCES "public"."homework"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "homework_submissions" ADD CONSTRAINT "homework_submissions_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "homework_submissions" ADD CONSTRAINT "homework_submissions_homework_id_homework_id_fk" FOREIGN KEY ("homework_id") REFERENCES "public"."homework"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "homework" ADD CONSTRAINT "homework_created_by_users_id_fk" FOREIGN KEY ("created_by") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "lesson_resources" ADD CONSTRAINT "lesson_resources_lesson_id_lessons_lesson_id_fk" FOREIGN KEY ("lesson_id") REFERENCES "public"."lessons"("lesson_id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "lesson_resources" ADD CONSTRAINT "lesson_resources_resource_id_resources_resource_id_fk" FOREIGN KEY ("resource_id") REFERENCES "public"."resources"("resource_id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "lessons" ADD CONSTRAINT "lessons_section_id_sections_section_id_fk" FOREIGN KEY ("section_id") REFERENCES "public"."sections"("section_id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "lessons" ADD CONSTRAINT "lessons_course_id_courses_courses_id_fk" FOREIGN KEY ("course_id") REFERENCES "public"."courses"("courses_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "notes" ADD CONSTRAINT "notes_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "notes" ADD CONSTRAINT "notes_lesson_id_lessons_lesson_id_fk" FOREIGN KEY ("lesson_id") REFERENCES "public"."lessons"("lesson_id") ON DELETE set null ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "organisation_users" ADD CONSTRAINT "organisation_users_organisation_id_organisations_organisations_id_fk" FOREIGN KEY ("organisation_id") REFERENCES "public"."organisations"("organisations_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "organisation_users" ADD CONSTRAINT "organisation_users_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "organisation_users" ADD CONSTRAINT "organisation_users_role_id_roles_role_id_fk" FOREIGN KEY ("role_id") REFERENCES "public"."roles"("role_id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "quizzes_questions" ADD CONSTRAINT "quizzes_questions_quizzes_id_quizzes_id_fk" FOREIGN KEY ("quizzes_id") REFERENCES "public"."quizzes"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "sections" ADD CONSTRAINT "sections_courses_id_courses_courses_id_fk" FOREIGN KEY ("courses_id") REFERENCES "public"."courses"("courses_id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "session" ADD CONSTRAINT "session_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE no action ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "user_roles" ADD CONSTRAINT "user_roles_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "user_roles" ADD CONSTRAINT "user_roles_role_id_roles_role_id_fk" FOREIGN KEY ("role_id") REFERENCES "public"."roles"("role_id") ON DELETE cascade ON UPDATE no action;--> statement-breakpoint
ALTER TABLE "weekly_planner" ADD CONSTRAINT "weekly_planner_class_id_classes_class_id_fk" FOREIGN KEY ("class_id") REFERENCES "public"."classes"("class_id") ON DELETE no action ON UPDATE no action;