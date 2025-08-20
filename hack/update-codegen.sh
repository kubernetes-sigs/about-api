# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

echo $CODEGEN_PKG

source "${CODEGEN_PKG}/kube_codegen.sh"

THIS_PKG="sigs.k8s.io/about-api"

kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate/boilerplate.generatego.txt" \
    "${SCRIPT_ROOT}/pkg"

kube::codegen::gen_register \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate/boilerplate.generatego.txt" \
    "${SCRIPT_ROOT}/pkg"

kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/pkg/generated" \
    --output-pkg "${THIS_PKG}/pkg/generated" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate/boilerplate.generatego.txt" \
    "${SCRIPT_ROOT}/pkg"
