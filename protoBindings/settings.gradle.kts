/*
 * This file was generated by the Gradle 'init' task.
 *
 * The settings file is used to specify which projects to include in your build.
 *
 * Detailed information about configuring a multi-project build in Gradle can be found
 * in the user manual at https://docs.gradle.org/7.2/userguide/multi_project_builds.html
 */

rootProject.name = "provenance"

/*
 * Dynamically include sub projects.
 * Uses the folder name as the project name
 */
File(rootDir, "bindings").walk().filter {
    it.isDirectory && File(it, "build.gradle.kts").isFile
}.forEach {
    include(it.name)
    project(":${it.name}").projectDir = it
}