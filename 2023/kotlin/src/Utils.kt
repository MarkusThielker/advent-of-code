import java.math.BigInteger
import java.security.MessageDigest
import kotlin.io.path.Path
import kotlin.io.path.readLines

fun readInput(name: String): List<String> {
    val path = System.getProperty("user.dir")
    return Path("$path/2023/kotlin/src/$name.txt").readLines()
}

fun String.md5() = BigInteger(1, MessageDigest.getInstance("MD5").digest(toByteArray()))
    .toString(16)
    .padStart(32, '0')

fun Any?.println() = println(this)
